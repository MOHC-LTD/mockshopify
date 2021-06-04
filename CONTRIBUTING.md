# How to contribute

## Contents

- [Git hooks](#-git-hooks)
- [Committing and merging](#-committing-and-merging)
- [Autogeneration](#-autogeneration)

## 🪝 Git hooks

This project includes two git hooks

- `pre-commit`
- `pre-push`

To enable these, run the following command

```sh
git config core.hooksPath .githooks
```

> Note - on a unix machine (mac and linux) you may need to make both files executable

```sh
chmod +x .githooks/pre-commit
chmod +x .githooks/pre-push
```

## 💬 Committing and merging

We use [commitlint](https://commitlint.js.org/) to ensure that all commit messages meet a certain style and standard. You can find out more about conventional commits [here](https://www.conventionalcommits.org/).

When merging pull requests into the `development` branch on GitHub, use the `Squash and merge` option and make sure the squash commit message follows the style from commitlint and contains the pull request number. Use the normal `Merge pull request` for everything else including the `main` branch.

## 🏎️ Autogeneration

The mock implementations can be auto generated using the [gomock](https://github.com/golang/mock) [`mockgen`](https://github.com/golang/mock#running-mockgen) command.

Clone both the [mockshopify](https://github.com/MOHC-LTD/mockshopify) and [shopify](https://github.com/MOHC-LTD/shopify) repositories onto your machine.

From the common parent directory run the following.

```sh
mockgen -source=shopify/orders.go -destination=mockshopify/orders.go -package mockshopify
```

This is with the exception of the `shop.go` file which must be manually updated.
