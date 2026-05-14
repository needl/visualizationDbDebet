import { appState } from '../../../../shared/state/appState.js';
import { fetchResponseData } from '../../api/statsApi.js';

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
