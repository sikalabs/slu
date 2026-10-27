package ip

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/time_utils"
	"github.com/spf13/cobra"
)

var FlagImage string
var FlagShell string
var FlagNode string
var FlagNamespace string
var FlagDryRun bool
var FlagDetached bool
var FlagNoRoot bool
var FlagRestricted bool

type kdevSeccompProfile struct {
	Type string `json:"type"`
}

type kdevSecurityContext struct {
	RunAsUser      *int64              `json:"runAsUser,omitempty"`
	RunAsGroup     *int64              `json:"runAsGroup,omitempty"`
	RunAsNonRoot   *bool               `json:"runAsNonRoot,omitempty"`
	SeccompProfile *kdevSeccompProfile `json:"seccompProfile,omitempty"`
}

type kdevCapabilities struct {
	Drop []string `json:"drop,omitempty"`
}

type kdevContainerSecurityContext struct {
	AllowPrivilegeEscalation *bool             `json:"allowPrivilegeEscalation,omitempty"`
	Capabilities             *kdevCapabilities `json:"capabilities,omitempty"`
}

type kdevContainer struct {
	Name            string                        `json:"name"`
	SecurityContext *kdevContainerSecurityContext `json:"securityContext,omitempty"`
}

type kdevPodSpec struct {
	NodeName        string               `json:"nodeName,omitempty"`
	SecurityContext *kdevSecurityContext `json:"securityContext,omitempty"`
	Containers      []kdevContainer      `json:"containers,omitempty"`
}

type kdevOverrides struct {
	Spec kdevPodSpec `json:"spec"`
}

var Cmd = &cobra.Command{
	Use:   "kdev",
	Short: "Run sikalabs/dev in Kubernetes",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		if FlagRestricted {
			FlagNoRoot = true
		}

		podName := "dev-" + strconv.Itoa(time_utils.Unix())
		kubectlRunArgs := []string{
			podName,
			"--image", FlagImage,
		}

		if !FlagDetached {
			kubectlRunArgs = append(kubectlRunArgs, "--rm", "-ti")
		}

		if FlagNode != "" || FlagNoRoot {
			overrides := kdevOverrides{
				Spec: kdevPodSpec{
					NodeName: FlagNode,
				},
			}

			if FlagNoRoot {
				uid := int64(1000)
				gid := int64(1000)
				runAsNonRoot := true
				overrides.Spec.SecurityContext = &kdevSecurityContext{
					RunAsUser:    &uid,
					RunAsGroup:   &gid,
					RunAsNonRoot: &runAsNonRoot,
				}
			}

			if FlagRestricted {
				allowPrivilegeEscalation := false
				overrides.Spec.SecurityContext.SeccompProfile = &kdevSeccompProfile{
					Type: "RuntimeDefault",
				}
				overrides.Spec.Containers = []kdevContainer{
					{
						Name: podName,
						SecurityContext: &kdevContainerSecurityContext{
							AllowPrivilegeEscalation: &allowPrivilegeEscalation,
							Capabilities: &kdevCapabilities{
								Drop: []string{"ALL"},
							},
						},
					},
				}
			}

			overridesJson, _ := json.Marshal(overrides)
			kubectlRunArgs = append(
				kubectlRunArgs,
				"--overrides", string(overridesJson),
			)
			if FlagRestricted {
				// merge containers by name instead of replacing the list
				kubectlRunArgs = append(kubectlRunArgs, "--override-type", "strategic")
			}
		}

		if FlagNamespace != "" {
			kubectlRunArgs = append(kubectlRunArgs, "--namespace", FlagNamespace)
		}

		kubectlArgs := append([]string{"run"}, kubectlRunArgs...)

		if !FlagDetached {
			kubectlArgs = append(kubectlArgs, "--", FlagShell)
		} else {
			kubectlArgs = append(kubectlArgs, "--", "sleep", "infinity")
		}

		if FlagDryRun {
			fmt.Printf(
				"kubectl %s\n",
				strings.Join(kubectlArgs, " "),
			)
		} else {
			cmd := exec.Command(
				"kubectl", kubectlArgs...,
			)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Run()
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagImage,
		"image",
		"i",
		"ghcr.io/sikalabs/dev",
		"Container Image",
	)
	Cmd.Flags().StringVarP(
		&FlagShell,
		"shell",
		"s",
		"zsh",
		"Shell to run in container",
	)
	Cmd.Flags().StringVar(
		&FlagNode,
		"node",
		"",
		"Node to run on",
	)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes namespace",
	)
	Cmd.Flags().BoolVarP(
		&FlagDetached,
		"detached",
		"d",
		false,
		"Run in detached mode",
	)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"print command instead of running it",
	)
	Cmd.Flags().BoolVar(
		&FlagNoRoot,
		"no-root",
		false,
		"Run container as non-root user (uid/gid 1000)",
	)
	Cmd.Flags().BoolVar(
		&FlagRestricted,
		"restricted",
		false,
		"Run with securityContext compliant with PodSecurity \"restricted\" (implies --no-root)",
	)
}
