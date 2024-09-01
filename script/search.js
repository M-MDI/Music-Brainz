function performSearch(term) {
    fetch(`/artist/{id}`)
        .then(response => response.json())
        .then(data => {
            displayResults(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function displayResults(results) {
    const resultsContainer = document.getElementById('search-results');
    resultsContainer.innerHTML = '';
    
    results.forEach(item => {
        const itemElement = document.createElement('div');
        itemElement.textContent = `${item.name}: ${item.description}`;
        resultsContainer.appendChild(itemElement);
    });
}