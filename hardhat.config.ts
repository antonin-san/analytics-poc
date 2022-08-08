import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.9",
  networks: {
    hardhat: {
      mining: {
        auto: false,
        interval: 0,
        mempool: {
          order: "fifo",
        },
      },
    },
  },
};

export default config;
