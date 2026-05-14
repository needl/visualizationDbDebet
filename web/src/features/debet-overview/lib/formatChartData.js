export function formatChartData(aggregated, metric) {
    const names = [];
    const values = [];

    for (const [name, data] of Object.entries(aggregated)) {
        names.push(name);
        values.push(data[metric] || 0);
    }

    return { names, values };
}
