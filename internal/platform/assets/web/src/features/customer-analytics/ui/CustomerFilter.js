import { appState } from '../../../shared/state/appState.js';
import { fetchCustomers } from '../api/customerApi.js';
import { loadCustomerData } from '../model/actions/loadCustomerData.js';

export class CustomerFilter {
    constructor(container) {
        this.container = container;
        this.select = null;
        this.init();
    }

    async init() {
        await this.loadCustomers();
        this.render();
        this.setupListeners();
    }

    async loadCustomers() {
        try {
            const customers = await fetchCustomers();
            appState.setCustomers(customers);
        } catch (err) {
            console.error('Failed to load customers', err);
        }
    }

    render() {
        const customers = appState._getState().customers;
        if (!customers.length) {
            this.container.innerHTML = '<div>Загрузка заказчиков...</div>';
            return;
        }

        const filteredCustomers = customers.filter((c) => c.name !== 'ГКУ УПТ');
        this.container.innerHTML = `
            <div class="customer-filter">
                <label for="customer-select">Выберите заказчика:</label>
                <select id="customer-select">
                    <option value="">-- Заказчик --</option>
                    ${filteredCustomers.map((c) => `<option value="${c.name}">${c.name}</option>`).join('')}
                </select>
            </div>
        `;
        this.select = this.container.querySelector('#customer-select');
    }

    setupListeners() {
        if (!this.select) return;

        this.select.addEventListener('change', (e) => {
            const orgName = e.target.value;
            if (!orgName) {
                appState.setSelectedCustomer(null);
                appState.setCustomerSummary(null);
                appState.setCustomerTopDebtors([]);
                appState.setCustomerTopOverdue([]);
                appState.setCustomerBlockFactors(null);
                appState.setCustomerLoading(false);
                appState.setCustomerError(null);
                return;
            }

            appState.setSelectedCustomer(orgName);
            loadCustomerData(orgName);
        });
    }
}
