// public/javascript/app.js
document.addEventListener('DOMContentLoaded', function () {
    // Initial fetch of recipes
    fetchRecipes();

    // Event listeners for filter buttons
    const filterButtons = document.querySelectorAll('.filter-btn');
    filterButtons.forEach(btn => {
        btn.addEventListener('click', function () {
            const status = this.getAttribute('data-status');
            filterRecipes(status);
        });
    });

    // Event listener for the search form
    document.getElementById('searchForm').addEventListener('submit', function (e) {
        e.preventDefault();
        const searchTerm = document.getElementById('searchInput').value;
        searchRecipes(searchTerm);
    });
});

// Fetch and display all recipes
function fetchRecipes() {
    fetch('/api/recipes')
        .then(response => response.json())
        .then(recipes => {
            displayRecipes(recipes);
        })
        .catch(error => console.error('Error:', error));
}

// Filter recipes by status
function filterRecipes(status) {
    fetch(`/api/recipes/status/${status}`)
        .then(response => response.json())
        .then(recipes => {
            displayRecipes(recipes);
        })
        .catch(error => console.error('Error:', error));
}

// Display recipes in the DOM
function displayRecipes(recipes) {
    const recipeListContainer = document.getElementById('recipeList');
    recipeListContainer.innerHTML = ''; // Clear existing recipes

    recipes.forEach(recipe => {
        const recipeElement = document.createElement('div');
        recipeElement.className = 'recipe-list-container';
        recipeElement.dataset.recipeID = recipe.id;
        recipeElement.innerHTML = `
            <h3>${recipe.title}</h3>
            <p>Status: ${recipe.status}</p>
            <p>Ingredients: ${recipe.ingredients}</p>
            <p>Instructions: ${recipe.instructions}</p>
        `;
        recipeListContainer.appendChild(recipeElement);
    });
}
function searchRecipes(searchTerm) {
    console.log('Searching for:', searchTerm); 
    fetch(`/search-recipes?term=${encodeURIComponent(searchTerm)}`)
        .then(response => response.json())
        .then(data => {
            console.log('Search results:', data); 
            updateRecipeList(data);
        })
        .catch(error => console.error('Error:', error));
}

function updateRecipeList(recipes) {
    const recipeListContainer = document.getElementById('recipeList');
    recipeListContainer.innerHTML = ''; // Clear previous results
    recipes.forEach(recipe => {
        // Adjust these according to your actual response structure
        const recipeElement = document.createElement('div');
        recipeElement.className = 'recipe-list-container';
        recipeElement.dataset.recipeID = recipe.id;
        recipeElement.innerHTML = `<p>${recipe.title} - ${recipe.status}</p>`;
        recipeListContainer.appendChild(recipeElement);
    });
}

   
    function updateRecipeStatus(recipeID, status) {
        fetch(`/api/recipes/${recipeID}/${status}`, { method: "POST" })
            .then(response => {
                if (response.ok) {
                    console.log(`Recipe status updated to ${status}`);
                } else {
                    console.error("Failed to update recipe status");
                }
            })
            .catch(error => console.error('Error:', error));
    }

    function deleteRecipe(recipeID) {
        fetch(`/api/recipes/${recipeID}`, {
            method: 'DELETE'
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.text();
            })
            .then(() => {
                fetchRecipes();
            })
            .catch(error => console.error('Error:', error));
    }

    const newForm = document.getElementById('new-form');

    newForm.addEventListener('submit', function (event) {
        event.preventDefault();
    
        const formData = new FormData(newForm);
        const jsonData = {};
    
        formData.forEach((value, key) => {
            jsonData[key] = value;
        });
    
        fetch('/api/recipes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(jsonData),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                console.log('Recipe added successfully:', data);
                window.location.href = '/recipes';
            })
            .catch(error => console.error('Error:', error));
    });

document.getElementById('searchForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const searchTerm = document.getElementById('searchInput').value;
    console.log('Searching for:', searchTerm); // Log the search term for debugging
    fetch(`/search-recipes?term=${encodeURIComponent(searchTerm)}`)
        .then(response => response.json())
        .then(data => {
            console.log('Search results:', data); // Log the results for debugging
            const resultsContainer = document.getElementById('searchResults');
            resultsContainer.innerHTML = ''; // Clear previous results
            data.forEach(recipe => {
                // Adjust these according to your actual response structure
                const recipeElement = document.createElement('div');
                recipeElement.innerHTML = `<h3>${recipe.title}</h3><p>${recipe.ingredients}</p><p>${recipe.instructions}</p>`;
                resultsContainer.appendChild(recipeElement);
            });
        })
        .catch(error => console.error('Error:', error));
});