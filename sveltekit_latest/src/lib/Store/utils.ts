function getCookieValue(cookieName: string) {
	// Split the cookies into an array

	if (typeof document !== 'undefined') {
		const cookies: string[] = document.cookie.split(';');

		// Loop through the array to find the cookie with the given name
		for (let i = 0; i < cookies.length; i++) {
			const cookie: string = cookies[i].trim();

			// Check if the cookie name matches the given name
			if (cookie.indexOf(cookieName + '=') == 0) {
				// Return the cookie value
				return cookie.substring(cookieName.length + 1, cookie.length);
			}
		}
	}

	// If the cookie is not found, return null
	return '';
}

export { getCookieValue };
