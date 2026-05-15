import { fetchContractorsByFactor } from '../../api/contractorApi.js';

export async function loadContractorsByFactor(orgName, columnName) {
    try {
        const data = await fetchContractorsByFactor(orgName, columnName);
        return data;
    } catch (err) {
        console.error('Failed to load contractors by factor:', err);
        throw err;
    }
}
