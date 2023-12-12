// public/javascript/app.js
document.addEventListener('DOMContentLoaded', function () {
    fetchRecipes();
    setupFilterButtons();
    setupSearchForm();
    setupRecipeActions();
    setupNewRecipeForm();
});

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


// Fetch and display all recipes
function fetchRecipes() {
fetch('/')
    .then(response => response.json())
    .then(recipes => {
        displayRecipes(recipes);
    })
    .catch(error => console.error('Error:', error));
}

function setupFilterButtons() {
    // Attach event listeners to filter buttons for recipes
    const filterCooking = document.getElementById("fltrCooking");
    filterCooking.addEventListener("click", () => filterRecipes('cooking'));

    const filterToCook = document.getElementById("fltrToCook");
    filterToCook.addEventListener("click", () => filterRecipes('to-cook'));

    const filterTried = document.getElementById("fltrTried");
    filterTried.addEventListener("click", () => filterRecipes('tried'));

    const filterNotTried = document.getElementById("fltrNotTried");
    filterNotTried.addEventListener("click", () => filterRecipes('not-tried'));
}


function setupSearchForm() {
    const searchForm = document.getElementById('searchForm');
    searchForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const searchTerm = document.getElementById('searchInput').value;
        searchRecipes(searchTerm);
    });
}

function filterRecipes(status) {
    fetch(`/api/recipes/status/${status}`)
        .then(response => response.json())
        .then(recipes => displayRecipes(recipes))
        .catch(error => console.error('Error filtering recipes:', error));
}

function searchRecipes(searchTerm) {
    console.log('Searching for:', searchTerm);
    fetch(`/api/recipes/search?query=${encodeURIComponent(searchTerm)}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Search results:', data);
            displaySearchResults(data);
        })
        .catch(error => console.error('Error:', error));
}


function displaySearchResults(recipes) {
    // Select the container where search results should be displayed
    const resultsContainer = document.getElementById('searchResults');
    // Clear any existing content
    resultsContainer.innerHTML = '';

    // Loop through the array of recipes
    recipes.forEach(recipe => {
        // Create a new div for each recipe
        const recipeElement = document.createElement('div');
        // Set the inner HTML of the div to display recipe details
        recipeElement.innerHTML = `
            <h3>${recipe.title}</h3>
            <p>${recipe.ingredients}</p>
            <p>${recipe.instructions}</p>
        `;
        // Append the new div to the container
        resultsContainer.appendChild(recipeElement);
    });
}

function displayRecipes(recipes) {
    const recipeListContainer = document.getElementById('recipeList');
    recipeListContainer.innerHTML = '';

    recipes.forEach(recipe => {
        const recipeElement = document.createElement('div');
        recipeElement.className = 'recipe-list-container';
        recipeElement.dataset.recipeID = recipe.id;
        recipeElement.innerHTML = `
            <h3>${recipe.title}</h3>
            <p>Status: ${recipe.status}</p>
            <p>Ingredients: ${recipe.ingredients}</p>
            <p>Instructions: ${recipe.instructions}</p>
            <button class="mark-as-button" data-status="cooking">Mark as Cooking</button>
            <button class="mark-as-button" data-status="to-cook">Mark as To Cook</button>
            <button class="mark-as-button" data-status="tried">Mark as Tried</button>
            <button class="delete-button">Delete</button>
        `;
        recipeListContainer.appendChild(recipeElement);
    });
}

function setupRecipeActions() {
    const recipeListContainer = document.getElementById('recipeList');
    recipeListContainer.addEventListener('click', function (e) {
        const recipeElement = e.target.closest('.recipe-list-container');
        if (!recipeElement) return;

        const recipeID = recipeElement.dataset.recipeID;
        if (e.target.matches('.mark-as-button')) {
            const status = e.target.dataset.status;
            updateRecipeStatus(recipeID, status);
        } else if (e.target.matches('.delete-button')) {
            deleteRecipe(recipeID);
        }
    });
}

function updateRecipeStatus(recipeID, newStatus) {
    // The rest of your updateRecipeStatus function remains unchanged.
}

function deleteRecipe(recipeID) {
    // The rest of your deleteRecipe function remains unchanged.
}

function setupNewRecipeForm() {
    const newForm = document.getElementById('new-form');
    newForm.addEventListener('submit', function (event) {
        event.preventDefault();
        const formData = new FormData(newForm);
        const jsonData = {};
    
        formData.forEach((value, key) => {
            jsonData[key] = value;
        });
    
        fetch('/recipe', {
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
            window.location.href = '/recipes'; // Update this if necessary to redirect to the correct URL
        })
        .catch(error => console.error('Error:', error));
    });
}
