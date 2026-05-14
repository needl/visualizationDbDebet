const HTTP_STATUS_MESSAGES = {
    400: 'некорректные параметры запроса.',
    401: 'требуется авторизация.',
    403: 'доступ запрещен.',
    404: 'данные не найдены.',
    408: 'сервер не ответил вовремя.',
    409: 'конфликт данных на сервере.',
    422: 'данные не прошли проверку.',
    429: 'слишком много запросов, попробуйте позже.',
    500: 'внутренняя ошибка сервера.',
    502: 'сервер временно недоступен.',
    503: 'сервер перегружен или на обслуживании.',
    504: 'истекло время ожидания ответа от сервера.'
};

function extractHttpStatus(error) {
    if (!error) return null;
    if (typeof error.status === 'number') return error.status;

    const message = String(error.message || '');
    const match = message.match(/\b(?:HTTP(?:\s+error!?[\s:]+)?)(\d{3})\b/i)
        || message.match(/\bstatus[:\s]+(\d{3})\b/i)
        || message.match(/\b(\d{3})\b/);

    if (!match) return null;
    const status = Number(match[1]);
    return Number.isInteger(status) ? status : null;
}

function isNetworkError(error) {
    const message = String(error?.message || '').toLowerCase();
    const name = String(error?.name || '').toLowerCase();
    return (
        name === 'typeerror'
        && (
            message.includes('failed to fetch')
            || message.includes('networkerror')
            || message.includes('network request failed')
            || message.includes('load failed')
        )
    );
}

function toSentence(text) {
    const value = String(text || '').trim();
    if (!value) return '';
    if (/[.!?]$/.test(value)) return value;
    return `${value}.`;
}

export function getUserFriendlyError(error, actionLabel = 'Не удалось загрузить данные') {
    if (isNetworkError(error)) {
        return `${actionLabel}: нет соединения с сервером. Проверьте, что API запущен и доступен.`;
    }

    const status = extractHttpStatus(error);
    if (status) {
        const statusMessage = HTTP_STATUS_MESSAGES[status] || `сервер вернул ошибку (${status}).`;
        return `${actionLabel}: ${statusMessage}`;
    }

    const original = String(error?.message || '').trim();
    if (!original) {
        return `${actionLabel}.`;
    }

    return `${actionLabel}: ${toSentence(original)}`;
}
