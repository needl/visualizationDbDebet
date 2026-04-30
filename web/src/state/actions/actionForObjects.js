// src/state/actions/actionForObjects.js
import { appState } from '../appState.js';
import { fetchObjectList, fetchObjectData } from '../../services/objectApiCaller.js';
import {aggregateObjectMetrics, prepareChartData} from "../../transformers/objectAggregator.js";

export async function loadObjectList(orgName) {
    appState.setObjectLoading(true);
    appState.setObjectError(null);
    try {
        const list = await fetchObjectList(orgName);
        appState.setObjectList(list);
        appState.setSelectedObject(null);
        appState.setObjectData(null);
    } catch (err) {
        console.error('Ошибка загрузки списка объектов:', err);
        appState.setObjectError(err.message);
        appState.setObjectList([]);
    } finally {
        appState.setObjectLoading(false);
    }
}

export async function loadObjectData(orgName, objectName) {
    appState.setObjectLoading(true);
    appState.setObjectError(null);
    try {
        const rawData = await fetchObjectData(orgName, objectName);
        const metrics = aggregateObjectMetrics(rawData);
        const chartData = prepareChartData(rawData);
        appState.setObjectData({ metrics, chartData });
    } catch (err) {
        console.error('Ошибка загрузки данных объекта:', err);
        appState.setObjectError(err.message);
        appState.setObjectData(null);
    } finally {
        appState.setObjectLoading(false);
    }
}