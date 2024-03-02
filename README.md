# BlogChain

**blog** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

### Built with

<div align="center">
    <img src="https://skillicons.dev/icons?i=go"/> <br>
</div>

## Get started

```
ignite chain serve
```

find your daemon file in the your compiled go files
put it in the cmd/blogd and called it
then you can call queries and do actions on the chain

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project.

For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

<hr>
Developer note:
Lasted update: 3/2/2024
- This chain is for studying purpose, not recommend to be  used as mainnet or testnet.
-  The codebase may have bugs that could potentially lead to loss of funds. Use at your own risk.
- Functions and queries is successfully tested, the repository does not include the daemon file, please follow the instruction or contact with the repo-owner for further assistance.
