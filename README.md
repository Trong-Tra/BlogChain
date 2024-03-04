# BlogChain

<h3 align="center">BlogChain</h3>
<p align="center">
A platform powered by blockchain that enables users to create and delete blogs. 
</p>

## About the projects

- Welcome to BlogChain, a blog application as a Cosmos SDK blockchain using the Ignite CLI.
- With BlogChain, you can create, read, update, and delete blog posts on the blockchain.

### Built with

<div align="center">
    <img src="https://skillicons.dev/icons?i=go"/> <br>
</div>

**BlogChain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

Run the chain with this command:

```
ignite chain serve
```

After the blockchain is running you need to have the daemon file of the chain in order for it to perform queries or call functions. <br>

First, find your daemon file in the your compiled go files. <br>
Put it in the cmd/blogd and called it. <br>
cd into the file (cmd/blogd) and from there you can perform queries or call functions. For example:

```
./blogd tx blog create-post hello world --from alice --chain-id blog
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project.

For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

<hr>
Developer note:
Lasted update: 3/2/2024 <br>

- This chain is for studying purpose, not recommend to be used as mainnet or testnet.
- The codebase may have bugs that could potentially lead to loss of funds. Use at your own risk.
- Functions and queries is successfully tested, the repository does not include the daemon file, please follow the instruction or contact with the repo-owner for further assistance.
