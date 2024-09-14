// theme-switcher.js

document.addEventListener('DOMContentLoaded', () => {
	const themeSwitcher = document.getElementById('themeSwitcher');
	const body = document.body;

	// Load the theme from localStorage
	const savedTheme = localStorage.getItem('theme') || 'light-mode';
	body.classList.add(savedTheme);

	// Set the initial icon
	themeSwitcher.textContent = savedTheme === 'light-mode' ? '🌙' : '🌞';

	themeSwitcher.addEventListener('click', () => {
		if (body.classList.contains('light-mode')) {
			body.classList.remove('light-mode');
			body.classList.add('dark-mode');
			themeSwitcher.textContent = '🌞'; // Sun icon for light mode
			localStorage.setItem('theme', 'dark-mode');
		} else {
			body.classList.remove('dark-mode');
			body.classList.add('light-mode');
			themeSwitcher.textContent = '🌙'; // Moon icon for dark mode
			localStorage.setItem('theme', 'light-mode');
		}
	});
});

