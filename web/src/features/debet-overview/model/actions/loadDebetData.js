import { appState } from '../../../../shared/state/appState.js';
import { fetchDebetData } from '../../api/debetApi.js';
import { aggregateByOrg } from '../../lib/aggregateByOrg.js';
import { formatChartData } from '../../lib/formatChartData.js';
import { prepareTableData } from '../../lib/prepareTableData.js';
import { aggregateByYearStacked } from '../../lib/aggregateByYearStacked.js';

export async function loadData() {
    appState.setLoading(true);
    try {
        const rawData = await fetchDebetData();
        const aggregated = aggregateByOrg(rawData);
        const contractData = formatChartData(aggregated, 'contractAmount');
        const debetTotalData = formatChartData(aggregated, 'debetTotal');
        const debetOverdoseData = formatChartData(aggregated, 'debetOverdose');
        const groupedYearData = aggregateByYearStacked(rawData);
        const tableData = prepareTableData(aggregated);

        const pieContract = {
            data: Object.entries(aggregated).map(([name, data]) => ({
                name,
                value: data.contractAmount || 0
            }))
        };

        const pieDebetTotal = {
            data: Object.entries(aggregated).map(([name, data]) => ({
                name,
                value: data.debetTotal || 0
            }))
        };

        const pieDebetOverdose = {
            data: Object.entries(aggregated).map(([name, data]) => ({
                name,
                value: data.debetOverdose || 0
            }))
        };

        appState.setChartData('contractAmount', contractData);
        appState.setChartData('debetTotal', debetTotalData);
        appState.setChartData('debetOverdose', debetOverdoseData);
        appState.setChartData('debetByYear', groupedYearData);
        appState.setChartData('contractAmountPie', pieContract);
        appState.setChartData('debetTotalPie', pieDebetTotal);
        appState.setChartData('debetOverdosePie', pieDebetOverdose);
        appState.setTableData(tableData);
        appState.setError(null);
        appState.setLoading(false);
    } catch (err) {
        appState.setError(err.message);
        appState.setLoading(false);
    }
}
