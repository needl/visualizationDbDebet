import { appState } from '../appState.js';
import {fetchResponseData} from "../../services/apiCaller.js";

export async function loadStats() {
    appState.setLoading(true);
    try {
        const stats = await fetchResponseData();
        appState.setStats(stats);
        appState.setError(null);
    } catch (err) {
        appState.setError(err.message);
    }
}