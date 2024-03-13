const State = artifacts.require("State");
const VerifierStateTransition = artifacts.require("VerifierStateTransition");
const PoseidonUnit1L = artifacts.require("PoseidonUnit1L");
const PoseidonUnit2L = artifacts.require("PoseidonUnit2L");
const PoseidonUnit3L = artifacts.require("PoseidonUnit3L");
const SmtLib = artifacts.require("SmtLib");
const StateLib = artifacts.require("StateLib");
const ERC1967Proxy = artifacts.require("ERC1967Proxy");

module.exports = async (deployer, logger) => {

    const { defaultIdType, chainId } = await getDefaultIdType();

    console.log(`found defaultIdType ${defaultIdType} for chainId ${chainId}`);

    deployer.startMigrationsBlock = await web3.eth.getBlockNumber();

    const verifier = await deployer.deploy(VerifierStateTransition);

    await deployer.deploy(PoseidonUnit1L);
    await deployer.deploy(PoseidonUnit2L);
    await deployer.deploy(PoseidonUnit3L);


    deployer.link(PoseidonUnit2L, SmtLib)
    deployer.link(PoseidonUnit3L, SmtLib)
    await deployer.deploy(SmtLib);

    await deployer.deploy(StateLib);

    deployer.link(PoseidonUnit1L, State)
    deployer.link(SmtLib, State)
    deployer.link(StateLib, State)
    const state  = await deployer.deploy(State);
    const proxy = await deployer.deploy(ERC1967Proxy, state.address, "0x");
    const stateProxy = await State.at(proxy.address);
    await State.setAsDeployed(stateProxy);

    logger.logTransaction(await stateProxy.initialize(verifier.address, defaultIdType), "Init State");
}

const chainIdDefaultIdTypeMap = new Map()
    .set(31337, '0x0212') // hardhat
    .set(80001, '0x0212') // polygon mumbai
    .set(1101,  '0x0231') // zkEVM
    .set(1442,  '0x0232') // zkEVM testnet
    .set(137,   '0x0211'); // polygon main

async function getDefaultIdType() {
    const chainId = parseInt(await network.provider.send('eth_chainId'), 16);
    const defaultIdType = chainIdDefaultIdTypeMap.get(chainId);
    if (!defaultIdType) {
        throw new Error(`Failed to find defaultIdType in Map for chainId ${chainId}`);
    }
    return { defaultIdType, chainId };
}
