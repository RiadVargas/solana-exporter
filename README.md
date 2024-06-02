# Solana Prometheus Exporter

The Solana Validator Prometheus Exporter is a tool designed to monitor and analyze the performance and health of your Solana validator node. By exporting detailed metrics to Prometheus, this exporter helps ensure your validator operates optimally, efficiently, and reliably.

## Content

<!-- TOC -->

- [Solana Prometheus Exporter](#solana-prometheus-exporter)
    - [Content](#content)
    - [How to use](#how-to-use)
        - [Docker](#docker)
    - [Configuration](#configuration)
        - [Arguments](#arguments)
        - [Options](#options)
        - [Complete command](#complete-command)
    - [Metrics](#metrics)
    - [Planned development](#planned-development)

<!-- /TOC -->

## How to use

### Docker

The easiest way to deploy this Solana exporter is using Docker: 

```
docker run -d --name solana-exporter --restart=on-failure -p 8888:8888 ghcr.io/riadvargas/solana-exporter <vote-account>
```

Replace `<vote-account>` with your validator vote account public address.

## Configuration

For smoother configuration and to avoid volume mounting, we rely on command params to configure the exporter:

### Arguments

- **voteKey** (a.k.a. vote account)
  - **Description**: The Solana validator vote account key to narrow resource-intensive searches to a specific validator. This argument is required for the exporter to function.

### Options

- **listenPort**
  - **Flag**: `--listenPort`
  - **Short**: `-p`
  - **Default**: `8888`
  - **Description**: Specifies the port on which the exporter will listen for incoming HTTP requests. This is the port that Prometheus will scrape to collect metrics from the exporter.

- **rpcUrl**
  - **Flag**: `--rpcUrl`
  - **Short**: `-r`
  - **Default**: `https://api.testnet.solana.com/`
  - **Description**: Specifies the Solana RPC URL that the exporter will connect to in order to gather metrics from the Solana network. Ensure that this URL points to a valid Solana RPC endpoint.

- **scrapeInterval**
  - **Flag**: `--scrapeInterval`
  - **Short**: `-i`
  - **Default**: `25 seconds`
  - **Description**: Defines the interval, in seconds, at which the exporter will scrape metrics from the Solana RPC. This controls how frequently data is collected and updated.

### Complete command

```
docker run -d --name solana-exporter -p 8888:8888 ghcr.io/riadvargas/solana-exporter --listenPort <port> --rpcUrl <rpc_url> --scrapeInterval <interval_in_seconds> <vote-account>
```

## Metrics

- **solana_current_epoch**
  - **Description**: Reports the current epoch

- **solana_first_slot**
  - **Labels**: `epoch` (The epoch number)
  - **Description**: Reports the first slot height for the epoch

- **solana_last_slot**
  - **Labels**: `epoch` (The epoch number)
  - **Description**: Reports the last expected slot for the epoch

- **solana_current_slot_height**
  - **Labels**: `epoch` (The epoch number)
  - **Description**: Reports the latest slot height for the epoch

- **solana_leadership_slot**
  - **Labels**: `epoch` (The epoch number), `slot` (The specific slot number), `identity` (The identity of the validator)
  - **Description**: Binary metric to foresee leadership slots for identities

- **solana_identity_balance**
  - **Labels**: `epoch` (The epoch number), `identity` (The identity of the validator), `vote_key` (The vote account key)
  - **Description**: Latest identity account balance for the epoch

- **solana_identity_last_vote_slot**
  - **Labels**: `epoch` (The epoch number), `identity` (The identity of the validator), `vote_key` (The vote account key)
  - **Description**: Last slot voted by the identity

- **solana_vote_epoch_credits**
  - **Labels**: `epoch` (The epoch number), `identity` (The identity of the validator), `vote_key` (The vote account key)
  - **Description**: Epoch credits for a given identity and vote key

- **solana_active_stake**
  - **Labels**: `epoch` (The epoch number), `identity` (The identity of the validator), `vote_key` (The vote account key)
  - **Description**: Activate stake for a given identity and vote key

- **solana_pending_activation_stake**
  - **Labels**: `epoch` (The epoch number), `identity` (The identity of the validator), `vote_key` (The vote account key)
  - **Description**: Pending for activation for a given identity and vote key

- **solana_validator_delinquent**
  - **Labels**: `epoch` (The epoch number), `identity` (The identity of the validator), `vote_key` (The vote account key)
  - **Description**: Binary metric that states if a vote account is delinquent

## Planned development
- [ ] Unit test metrics calculation
- [ ] Add a Grafana dashboard template
- [ ] New metric: cluster average vote credits
- [ ] New metric: skipped slots
- [ ] Optimize once per epoch calculations