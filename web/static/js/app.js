function getCredentials() {
    const idInstance = document.getElementById('idInstance').value.trim();
    const apiTokenInstance = document.getElementById('apiTokenInstance').value.trim();

    if (!idInstance || !apiTokenInstance) {
        appendResponse('Error: Please enter idInstance and ApiTokenInstance');
        return null;
    }

    return { idInstance, apiTokenInstance };
}

function appendResponse(text) {
    const responseArea = document.getElementById('response');
    const timestamp = new Date().toLocaleTimeString();
    responseArea.value += `[${timestamp}] ${text}\n\n`;
    responseArea.scrollTop = responseArea.scrollHeight;
}

function clearResponse() {
    document.getElementById('response').value = '';
}

async function makeRequest(endpoint, body) {
    try {
        const response = await fetch(`/api/${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        });

        const data = await response.json();

        if (data.success) {
            return data.data;
        } else {
            throw new Error(data.error || 'Unknown error');
        }
    } catch (error) {
        throw error;
    }
}

async function getSettings() {
    const credentials = getCredentials();
    if (!credentials) return;

    appendResponse('Calling getSettings...');

    try {
        const result = await makeRequest('getSettings', credentials);
        appendResponse('getSettings response:\n' + JSON.stringify(result, null, 2));
    } catch (error) {
        appendResponse('getSettings error: ' + error.message);
    }
}

async function getStateInstance() {
    const credentials = getCredentials();
    if (!credentials) return;

    appendResponse('Calling getStateInstance...');

    try {
        const result = await makeRequest('getStateInstance', credentials);
        appendResponse('getStateInstance response:\n' + JSON.stringify(result, null, 2));
    } catch (error) {
        appendResponse('getStateInstance error: ' + error.message);
    }
}

async function sendMessage() {
    const credentials = getCredentials();
    if (!credentials) return;

    const phoneNumber = document.getElementById('phoneNumber').value.trim();
    const message = document.getElementById('message').value.trim();

    if (!phoneNumber || !message) {
        appendResponse('Error: Please enter phone number and message');
        return;
    }

    const body = {
        ...credentials,
        phoneNumber,
        message
    };

    appendResponse('Calling sendMessage...');

    try {
        const result = await makeRequest('sendMessage', body);
        appendResponse('sendMessage response:\n' + JSON.stringify(result, null, 2));
    } catch (error) {
        appendResponse('sendMessage error: ' + error.message);
    }
}

async function sendFileByUrl() {
    const credentials = getCredentials();
    if (!credentials) return;

    const phoneNumber = document.getElementById('filePhoneNumber').value.trim();
    const fileUrl = document.getElementById('fileUrl').value.trim();
    const fileName = document.getElementById('fileName').value.trim();
    const caption = document.getElementById('caption').value.trim();

    if (!phoneNumber || !fileUrl || !fileName) {
        appendResponse('Error: Please enter phone number, file URL, and file name');
        return;
    }

    const body = {
        ...credentials,
        phoneNumber,
        fileUrl,
        fileName,
        caption
    };

    appendResponse('Calling sendFileByUrl...');

    try {
        const result = await makeRequest('sendFileByUrl', body);
        appendResponse('sendFileByUrl response:\n' + JSON.stringify(result, null, 2));
    } catch (error) {
        appendResponse('sendFileByUrl error: ' + error.message);
    }
}
