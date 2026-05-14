import { dashboardConfig } from './dashboard/config.js';
import { loadData } from '../features/debet-overview/model/actions/loadDebetData.js';
import { loadStats } from '../features/debet-overview/model/actions/loadStats.js';
import { DashboardRenderer } from './dashboard/renderer.js';

const renderer = new DashboardRenderer(dashboardConfig);
const app = document.getElementById('app');

if (app) {
    renderer.build(app);
}

loadStats();
loadData();
