import { network } from "hardhat";

async function main() {
  await network.provider.send("hardhat_mine");
  console.log("Block mined");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
