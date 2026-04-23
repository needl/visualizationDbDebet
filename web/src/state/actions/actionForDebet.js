// src/state/actionForDebet.js
import { appState } from '../appState.js';
import { fetchDebetData } from '../../services/apiCaller.js';
import { aggregateByOrg } from '../../transformers/aggregatorByOrg.js';
import { formatChartData } from '../../transformers/formatter.js';
//import { aggregateByYear } from '../../transformers/groupByYear.js';
import { prepareTableData } from '../../transformers/tableDate.js';
import {aggregateByYearStacked} from "../../transformers/groupByYear.js";

export async function loadData() {
    appState.setLoading(true);
    try {
        const rawData = await fetchDebetData();
        const aggregated = aggregateByOrg(rawData);
        const contractData = formatChartData(aggregated, 'contractAmount');
        const debetTotalData = formatChartData(aggregated, 'debetTotal');
        const debetOverdoseData = formatChartData(aggregated, 'debetOverdose');

        //const groupedYearData = aggregateByYear(rawData);
        const groupedYearData = aggregateByYearStacked(rawData);
        const tableData = prepareTableData(aggregated);

        // Данные для круговых диаграмм по трём метрикам
        const pieContract = {
            data: Object.entries(aggregated).map(([name, data]) => ({
                name: name,
                value: data.contractAmount || 0
            }))
        };
        const pieDebetTotal = {
            data: Object.entries(aggregated).map(([name, data]) => ({
                name: name,
                value: data.debetTotal || 0
            }))
        };
        const pieDebetOverdose = {
            data: Object.entries(aggregated).map(([name, data]) => ({
                name: name,
                value: data.debetOverdose || 0
            }))
        };

        appState.setChartData('contractAmount', contractData);
        appState.setChartData('debetTotal', debetTotalData);
        appState.setChartData('debetOverdose', debetOverdoseData);
        appState.setChartData('debetByYear', groupedYearData);

        // Сохраняем данные для круговых диаграмм
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