// src/state/appState.js
class AppState {
    constructor() {
        this._chartData = {
            contractAmount: null,
            debetTotal: null,
            debetOverdose: null,
            debetByYear: null
        };
        this._stats = null;
        this._tableData = null;
        this._loading = false;
        this._error = null;
        this._subscribers = [];
        this._customers = [];                   // список заказчиков для фильтра
        this._selectedCustomer = null;          // выбранный заказчик (name)
        this._customerSummary = null;           // summary для выбранного
        this._customerTopDebtors = [];          // топ‑10 по ДЗ
        this._customerTopOverdue = [];          // топ‑10 по ПДЗ
        this._customerBlockFactors = null;      // блок-факторы
        this._customerLoading = false;           // флаг загрузки данных заказчика
        this._customerError = null;

        this._objectList = [];             // массив строк – имена объектов
        this._selectedObject = null;      // выбранный объект (имя)
        this._objectData = null;          // { metrics: {...}, chartData: {...} }
        this._objectLoading = false;
        this._objectError = null;
    }

    setCustomers(list) { this._customers = list; this._notify(); }
    setSelectedCustomer(customer) { this._selectedCustomer = customer; this._notify(); }
    setCustomerSummary(data) {this._customerSummary = data; this._notify(); }
    setCustomerTopDebtors(data) { this._customerTopDebtors = data; this._notify(); }
    setCustomerTopOverdue(data) { this._customerTopOverdue = data; this._notify(); }
    setCustomerBlockFactors(data) { this._customerBlockFactors = data; this._notify(); }
    setCustomerLoading(loading) { this._customerLoading = loading; this._notify(); }
    setCustomerError(error) { this._customerError = error; this._notify(); }

    setObjectList(list) { this._objectList = list; this._notify(); }
    setSelectedObject(name) { this._selectedObject = name; this._notify(); }
    setObjectData(data) { this._objectData = data; this._notify(); }
    setObjectLoading(loading) { this._objectLoading = loading; this._notify(); }
    setObjectError(error) { this._objectError = error; this._notify(); }

    subscribe(callback) {
        this._subscribers.push(callback);
        callback(this._getState());
        return () => {
            const index = this._subscribers.indexOf(callback);
            if (index !== -1) this._subscribers.splice(index, 1);
        };
    }

    _notify() {
        const state = this._getState();
        this._subscribers.forEach(cb => cb(state));
    }

    setChartData(metric, data) {
        this._chartData[metric] = data;
        this._notify();
    }

    setLoading(loading) {
        this._loading = loading;
        this._notify();
    }

    setError(error) {
        this._error = error;
        this._loading = false;
        this._notify();
    }

    setStats(stats) {
        this._stats = stats;
        this._notify();
    }

    setTableData(data) {               // новый метод
        this._tableData = data;
        this._notify();
    }


    _getState() {
        return {
            chartData: this._chartData,
            stats: this._stats,
            tableData: this._tableData,
            loading: this._loading,
            error: this._error,
            customers: this._customers,
            selectedCustomer: this._selectedCustomer,
            customerSummary: this._customerSummary,
            customerTopDebtors: this._customerTopDebtors,
            customerTopOverdue: this._customerTopOverdue,
            customerBlockFactors: this._customerBlockFactors,
            customerLoading: this._customerLoading,
            customerError: this._customerError,

            objectList: this._objectList,
            selectedObject: this._selectedObject,
            objectData: this._objectData,
            objectLoading: this._objectLoading,
            objectError: this._objectError
        };
    }
}

export const appState = new AppState();