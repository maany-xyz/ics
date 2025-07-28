#!/bin/bash

# Prerequisits:
# -> must create build folder before (see README.md)

#TODO:
# - Replace all investor, advisor and team wallets with already created addresses
# - Exact days of lockup, vesting

# Configuration
PROVIDER_HOME="$HOME/.maany-ics-p"
GENESIS_PATH="${PROVIDER_HOME}/config/genesis.json"
BINARY="maany-provider"
CHAIN_ID="maany-mainnet"
TOKEN_DENOM="stake"

VALIDATOR_ACCOUNT=(
  "validator:50000000000000:5000000000000"
)

# Check if gensis file exist. Create if it doesn't exist.
if [ ! -f "$GENESIS_PATH" ]; then
  echo "Initializing node as genesis.json is missing..."
  $BINARY init "hub-node" --chain-id "$CHAIN_ID"
fi

echo "Gensis.json exists or was successfully created..."

# Function to get wallet address
get_wallet_address() {
  local wallet_name=$1
  $BINARY keys show "$wallet_name" --keyring-backend test -a
}

# Add validator account
echo "Adding validator account..."
for account in "${VALIDATOR_ACCOUNT[@]}"; do
  IFS=":" read -r wallet_name amount self_stake <<< "$account"
  
  sleep 5
  # Create the wallet (if it doesn't already exist)
  if ! $BINARY keys show "$wallet_name" --keyring-backend test &> /dev/null; then
    $BINARY keys add "$wallet_name" --keyring-backend test --output json > $PROVIDER_HOME/keypair_$wallet_name.json
    echo "Created wallet: $wallet_name"
  fi
  
  # Get the wallet address
  wallet_address=$(get_wallet_address "$wallet_name")
  echo "in here with wallet address ${wallet_address}"
  # Add to genesis
  $BINARY genesis add-genesis-account "$wallet_address" "${amount}${TOKEN_DENOM}" --keyring-backend test
  echo "Added $amount $TOKEN_DENOM to $wallet_name ($wallet_address)"

  # Add gentx entry to create Validator with self-delegation
  $BINARY genesis gentx "$wallet_name" "${self_stake}${TOKEN_DENOM}" --chain-id "$CHAIN_ID" --keyring-backend test
  echo "Created gentx-file successully."

  # Collect gentx information
  $BINARY genesis collect-gentxs
  echo "Collected gentx-file successully."

done
