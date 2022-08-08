import { ethers } from "hardhat";

async function main() {

  const Lock = await ethers.getContractFactory("Lock");
  const lock = Lock.attach("0x5fbdb2315678afecb367f032d93f642f64180aa3");

  let index
  const promises = []
  //Math.floor(Math.random() * 300)
  for (index = 0; index < 265; index++) {
     promises.push(lock.withdraw());
  }

  await Promise.all(promises);

  console.log("Transactions in next block:", index);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
