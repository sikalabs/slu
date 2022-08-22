fmt:
	terraform fmt -recursive

fmt-check:
	terraform fmt -recursive -check

fmt-check-diff:
	terraform fmt -recursive -check

setup-git-hooks:
	rm -rf .git/hooks
	(cd .git && ln -s ../.git-hooks hooks)
