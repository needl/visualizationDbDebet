// src/state/actionForCustomer.js
import { appState } from '../appState.js';
import {
    fetchCustomerSummary,
    fetchCustomerTopDebtors,
    fetchCustomerTopOverdue,
    fetchCustomerBlockFactors
} from '../../services/customerApiCaller.js';

export async function loadCustomerData(orgName) {
    if (!orgName) return;
    appState.setCustomerLoading(true);
    appState.setCustomerError(null);
    try {
        const [summary, topDebtors, topOverdue, blockFactors] = await Promise.all([
            fetchCustomerSummary(orgName),
            fetchCustomerTopDebtors(orgName),
            fetchCustomerTopOverdue(orgName),
            fetchCustomerBlockFactors(orgName)
        ]);
        console.log('summary in loadCustomerData:', summary);
        appState.setCustomerSummary(summary);
        appState.setCustomerTopDebtors(topDebtors);
        appState.setCustomerTopOverdue(topOverdue);
        appState.setCustomerBlockFactors(blockFactors);
    } catch (err) {
        appState.setCustomerError(err.message);
        console.error('Failed to load customer data', err);
    } finally {
        appState.setCustomerLoading(false);
    }
}

