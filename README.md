# Interchain Security

# Set-up

1.  clone repo
2.  run `make install`
    -> this creates the "maany-provider" cli
3.  initialize genesis.json:
    - `maany-provider init hub-node --chain-id maany-mainnet`
      -> this creates a folder in your root called `.maany-ics-p`
      -> the folder contains all the relevant config files to configure the chain before starting it
4.  Initialize gentx file and validator account with self delegation
    - Check if `init.gentx.sh` (in root of project) inputs and paths match your local setup
    - run `./init-gentx.sh` from the root of the ics repo
5.  In the genesis.json of the config folder add the module account under `app_state.auth.accounts`

        {
          "@type": "/cosmos.auth.v1beta1.ModuleAccount",
          "base_account": {
            "address": "maany1kdsm4jzhnrck2ucykhrj8lhhayp3am3sl9n4k3",
            "pub_key": null,
            "account_number": "1",
            "sequence": "0"
          },
          "name": "blockrewards",
          "permissions": ["minter"]
        }

    under `app_state.bank.balances` add

        {
          "address": "maany1kdsm4jzhnrck2ucykhrj8lhhayp3am3sl9n4k3",
          "coins": [
            {
              "denom": "stake",
              "amount": "300000000000000"
            }
          ]
        },

6.  Create one or more test accounts (makes sense to create 2 accounts and test transactions once the chain is running)

- `maany-provider keys add <account_name> --keyring-backend test`
  -> this creates a test account and keyring is stored directly in the config folder

6. Add an amount to a new account on chain start:

- `maany-provider genesis add-genesis-account <address> <amount, e.g. 100000stake> --keyring-backend test`

- Note: Double check in genesis.json, your account should've been added under `app_state.auth.accounts`
  {
  "@type": "/cosmos.auth.v1beta1.BaseAccount",
  "address": <your_address>,
  "pub_key": null,
  "account_number": "0",
  "sequence": "0"
  },
- then under `app_state.bank.balances` add funds to those accounts like by adding:
  {
  "address": <your_address>,
  "coins": [
  {
  "denom": "stake",
  "amount": <number e.g. 10000000000stake>
  }
  ]
  },

- IMPORTANT: the sum of all balances in `app_state.bank.balances` must the same as in `app_state.bank.supply.amount`
  e.g. "supply": [
  {
  "denom": "stake",
  "amount": "10000000000"
  }
  ],

7. Run the chain `maany-provider start`

# Commands

# Check Account balance

`maany-provider q bank balances <address>`

# Send transaction

`maany-provider tx bank send <from_key_or_address> [to_address] <amount>`
e.g. maany-provider tx bank send maany123... maany234... 1000000stake
