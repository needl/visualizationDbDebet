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
        this._tableData = null;   // новое поле
        this._loading = false;
        this._error = null;
        this._subscribers = [];
    }

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
            error: this._error
        };
    }
}

export const appState = new AppState();