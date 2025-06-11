const eventToast = "toast";

// Notify
function notify(title, message, type = 'info') {
    const notificationContainer = document.getElementById('notification-container');

    // Create notification element
    const notification = document.createElement('div');
    notification.className = `alert alert-${type} shadow-lg animate-fadeIn mb-2`;
    notification.innerHTML = `
                <div>
                    <i class="fas fa-${type === 'info' ? 'info-circle' : type === 'success' ? 'check-circle' : type === 'warning' ? 'exclamation-triangle' : 'times-circle'}"></i>
                    <div>
                        <h3 class="font-bold">${title}</h3>
                        <div class="text-xs">${message}</div>
                    </div>
                </div>
                <button class="btn btn-sm btn-ghost" onclick="this.parentElement.remove()">
                    <i class="fas fa-times"></i>
                </button>
            `;

    // Add notification to container
    notificationContainer.prepend(notification);

    // Auto-remove after 5 seconds
    setTimeout(() => {
        notification.classList.add('opacity-0');
        setTimeout(() => notification.remove(), 300);
    }, 5000);
}

// --- Cookie Toast Notification System ---

/**
 * Gets a cookie by its name.
 * @param {string} name - The name of the cookie.
 * @returns {string|null} The cookie value or null if not found.
 */
function getCookie(name) {
    const nameEQ = name + "=";
    const ca = document.cookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') c = c.substring(1, c.length);
        if (c.indexOf(nameEQ) === 0) return decodeURIComponent(c.substring(nameEQ.length, c.length));
    }
    return null;
}

/**
 * Deletes a cookie by its name.
 * Sets the cookie's expiration date to the past.
 * @param {string} name - The name of the cookie to delete.
 */
function deleteCookie(name) {
    // Set path to / to ensure it's deleted regardless of the path it was set with,
    // assuming it's a root-level or general purpose cookie.
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

/**
 * Checks for a 'toast' cookie, parses its JSON content,
 * dispatches a 'toast' custom event, and then deletes the cookie.
 */
function processToastCookieAndDispatchEvent() {
    const cookieName = 'toast';
    const cookieValue = atob(getCookie(cookieName));

    if (cookieValue) {
        try {
            const toastData = JSON.parse(cookieValue);
            if (toastData) {
                const event = new CustomEvent(eventToast, { detail: toastData })
                document.body.dispatchEvent(event);
            } else {
                console.warn('Toast cookie found, but data is incomplete or malformed. Expected {type, title, message}. Received:', toastData);
            }
        } catch (e) {
            console.error('Failed to parse toast cookie JSON:', cookieValue, e);
        } finally {
            // Always attempt to delete the cookie after processing
            deleteCookie(cookieName);
        }
    }
}

document.addEventListener("DOMContentLoaded", function () {
    document.body.addEventListener(eventToast, (e) => {
        notify(e.detail.title, e.detail.message, e.detail.type);
    });

    // Process any toast notification from a cookie on page load
    processToastCookieAndDispatchEvent();
});

// Copy API key to clipboard
function copyApiKey() {
    const apiKeyInput = document.getElementById('api-key');
    apiKeyInput.select();
    navigator.clipboard.writeText(apiKeyInput.value);

    notify('Success', 'API key copied to clipboard!', 'success');
}

// Add this function to /static/js/app.js
function switchTab(clickedTabElement) {
    // Find the parent '.tabs' container of the clicked tab
    const tabGroup = clickedTabElement.closest('.tabs');
    if (!tabGroup) {
        console.error('Tab group (.tabs) not found for clicked element:', clickedTabElement);
        return;
    }

    // Get all tab links within this specific tab group
    const tabLinks = tabGroup.querySelectorAll('.tab');

    // Get the main container for all tab content panels
    // Based on your HTML, this is <div id="tab-content">
    const tabContentContainer = document.getElementById('tab-content');
    if (!tabContentContainer) {
        console.error('Tab content container (#tab-content) not found.');
        return;
    }

    // Get all individual tab content panels.
    // These panels have IDs like "results-tab-content", "raw-tab-content", etc.
    const tabPanels = tabContentContainer.querySelectorAll('div[id$="-tab-content"]');

    // Deactivate all tabs (remove 'tab-active' class)
    tabLinks.forEach(tab => {
        tab.classList.remove('tab-active');
    });

    // Hide all tab content panels (add 'hidden' class)
    tabPanels.forEach(panel => {
        panel.classList.add('hidden');
    });

    // Activate the clicked tab
    clickedTabElement.classList.add('tab-active');

    // Show the corresponding content panel
    // The clicked tab has an ID like "results-tab".
    // The corresponding content panel ID is "results-tab-content".
    const contentPanelId = clickedTabElement.id + '-content';
    const activePanel = document.getElementById(contentPanelId);

    if (activePanel) {
        activePanel.classList.remove('hidden');
    } else {
        console.error('Could not find tab content panel with ID:', contentPanelId, 'for tab:', clickedTabElement.id);
    }
}
