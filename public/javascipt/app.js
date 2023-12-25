
document.addEventListener('DOMContentLoaded', function() {
    fetch('/api/recipes') // Adjusted endpoint
        .then(response => response.json())
        .then(recipes => {
            const selectElement = document.getElementById('recipe-select');
            recipes.forEach(recipe => {
                const option = new Option(recipe.title, recipe.ID); // Adjust to match your Recipe model
                selectElement.add(option);
            });
        })
        .catch(error => console.error('Error:', error));
});

document.addEventListener('DOMContentLoaded', function () {
    fetchRecipes();
    attachEventListeners();
    populateRecipeDropdown();
});

function attachEventListeners() {

        document.querySelectorAll('.filter-btn').forEach(button => {
            button.addEventListener('click', function() {
                const status = this.getAttribute('data-status');
                navigateToStatusPage(status); // Call a function to navigate to the status-specific page
            });
        });
    
    const recipeList = document.getElementById('recipeList');
    if (recipeList) {
        recipeList.addEventListener('click', handleRecipeListClick);
    }

    const newForm = document.getElementById('new-form');
    const updateForm = document.getElementById('update-form');
    const deleteForm = document.getElementById('delete-form');
    document.getElementById('delete-form').addEventListener('submit', function(event) {
        event.preventDefault();
        const recipeID = document.getElementById('recipe-id').value;
        deleteRecipe(recipeID);
    });
    

    if (newForm) {
        newForm.addEventListener('submit', handleNewRecipeSubmit);
    }
    if (updateForm) {
        updateForm.addEventListener('submit', handleFormSubmitWithFetch('/api/recipes/update', 'POST'));
    }
    if (deleteForm) {
        deleteForm.addEventListener('submit', handleFormSubmitWithFetch('/api/recipes/delete', 'POST'));
    }
}
document.getElementById('searchForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const query = document.getElementById('searchInput').value;
    const searchType = document.getElementById('searchType').value;
    redirectToSearchResults(query, searchType);
});

function redirectToSearchResults(query, searchType) {
    let url;
    switch(searchType) {
        case 'title':
            url = `/api/recipes/title/${encodeURIComponent(query)}`;
            break;
        case 'ingr':
            url = `/api/recipes/ingr/${encodeURIComponent(query)}`;
            break;
        case 'instr':
            url = `/api/recipes/instr/${encodeURIComponent(query)}`;
            break;
        default:
            console.error('Unknown search type');
            return;
    }
    window.location.href = url; // Перенаправление на новый URL
}



function searchRecipes(query, searchType) {
    let url;
    switch(searchType) {
        case 'title':
            url = `/api/recipes/title/${encodeURIComponent(query)}`;
            break;
        case 'ingr':
            url = `/api/recipes/ingr/${encodeURIComponent(query)}`;
            break;
        case 'instr':
            url = `/api/recipes/instr/${encodeURIComponent(query)}`;
            break;
        default:
            console.error('Unknown search type');
            return;
    }

    fetch(url)
        .then(response => response.json())
        .then(recipes => {
            displaySearchResults(recipes);
        })
        .catch(error => console.error('Error:', error));
}

function displaySearchResults(recipes) {
    const resultsContainer = document.getElementById('searchResults');
    resultsContainer.innerHTML = ''; 

    if (recipes.length === 0) {
        resultsContainer.innerHTML = '<p>No recipes found.</p>';
        return;
    }

    recipes.forEach(recipe => {
        const recipeElement = document.createElement('div');
        recipeElement.textContent = recipe.title;
        resultsContainer.appendChild(recipeElement);
    });
}

function handleRecipeListClick(event) {
    if (event.target.matches('.delete-button')) {
        const recipeID = event.target.dataset.recipeid;
        deleteRecipe(recipeID);
    }
    if (event.target.matches('.mark-as-button')) {
        const recipeID = event.target.dataset.recipeid;
        const status = event.target.dataset.status;
        updateRecipeStatus(recipeID, status);
    }
}
    function fetchRecipes() {
        // Fetch the initial HTML content
        fetch('/')
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.text(); // Parse the HTML content
            })
            .then(htmlContent => {
                // Create a temporary container to parse the HTML
                const tempContainer = document.createElement('div');
                tempContainer.innerHTML = htmlContent;

                // Extract the recipe list container from the parsed HTML
                const recipeListContainer = tempContainer.querySelector('#recipeList');

                if (recipeListContainer) {
                    // Update the existing recipe list container with the received HTML content
                    const currentRecipeListContainer = document.getElementById('recipeList');
                    currentRecipeListContainer.innerHTML = recipeListContainer.innerHTML;

                    // Add click event listeners to each recipe element
                    currentRecipeListContainer.addEventListener('click', (event) => {
                        const recipeElement = event.target.closest('.recipe-list-container');
                        if (recipeElement) {
                            const recipeID = recipeElement.dataset.recipeId;
                            showRecipeDetails(recipeID);
                        }
                    });

                } else {
                    console.error('Recipe list container not found in the received HTML.');
                }

            })
            .catch(error => console.error('Error:', error));
    }
    function filterRecipes(status) {
        var recipeListContainer = document.getElementById("recipeList");
        let url = `/api/${status}`;
        console.log("Fetching recipes with status:", status, "from", url); // Debugging line
    
        fetch(url)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(recipes => {
                console.log("Recipes received:", recipes); // Debugging line
                recipeListContainer.innerHTML = '';
                recipes.forEach(recipe => {
                    var recipeElement = document.createElement('div');
                    recipeElement.innerHTML = createRecipeHtml(recipe);
                    recipeListContainer.appendChild(recipeElement);
                });
            })
            .catch(error => console.error('Error:', error));
    }
    
    function navigateToStatusPage(status) {
        // Construct the URL for the status-specific page
        const url = `/${status}`;
    
        // Navigate to the new URL
        window.location.href = url;
    }
    // Handle click events to display recipe details
    function showRecipeDetails(recipeID) {
        fetch(`/api/recipes/${recipeID}`)
            .then(response => response.json())
            .then(recipe => {
                const recipeDetailsContainer = document.getElementById('recipeDetailsContainer');
                const titleElement = document.getElementById('recipeTitle');
                const ingredientsElement = document.getElementById('recipeIngredients');
                const instructionsElement = document.getElementById('recipeInstructions');
                const statusElement = document.getElementById('recipeStatus');
                const idElement = document.getElementById('recipeID');

                titleElement.textContent = recipe.title;
                ingredientsElement.textContent = `Ingredients: ${recipe.ingredients}`;
                instructionsElement.textContent = `Instructions: ${recipe.instructions}`;
                statusElement.textContent = `Status: ${recipe.status}`;
                idElement.textContent = recipe.id;

                // Show the details container
                recipeDetailsContainer.style.display = 'block';
            })
            .catch(error => console.error('Error fetching recipe details:', error));
    }

    // Function to update recipe status
 // Update recipe status function
function updateRecipeStatus(recipeID, newStatus) {
    // Assuming you're sending a POST request to update the status
    // This might vary depending on your API implementation
    fetch(`/api/recipes/${recipeID}/${newStatus}`, { method: 'POST' })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            console.log(`Recipe status updated to ${newStatus}`);
            // Here you might want to update the UI accordingly
        })
        .catch(error => console.error('Error:', error));
}

function populateRecipeDropdown() {
    fetch('/api/recipes/titles')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(recipes => {
            console.log(recipes); // Check the recipes in the console
            const selectElement = document.getElementById('recipe-select');
            recipes.forEach(recipe => {
                const option = new Option(recipe.title, recipe.id);
                selectElement.add(option);
            });
        })
        .catch(error => {
            console.error('Error fetching recipe titles:', error);
        });
}

document.getElementById('delete-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const recipeID = document.getElementById('recipe-select').value;
    deleteRecipe(recipeID);
});

function deleteRecipe(recipeID) {
    fetch(`/api/recipes/${recipeID}`, { method: 'DELETE' })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            console.log('Recipe deleted successfully');
            window.location.href = '/'; // Redirect to the updated list of recipes
        })
        .catch(error => console.error('Error:', error));
}


    // Function to update the recipe list on the page
    function updateRecipeList(recipes) {
        const recipeListContainer = document.getElementById('recipeList');
        recipeListContainer.innerHTML = '';

        if (recipes.length > 0) {
            recipes.forEach(recipe => {
                const recipeDiv = document.createElement('div');
                const recipeID = recipe.ID;

                recipeDiv.innerHTML = `
                    <p onclick="showRecipeDetails('${recipeID}')" style="cursor: pointer;">
                        ${recipe.title} by ${recipe.ingredients}: ${recipe.status}
                    </p>
                    <button class="mark-as-button" data-recipeid="${recipeID}" data-status="cooking">Cooking</button>
                    <button class="mark-as-button" data-recipeid="${recipeID}" data-status="tried">Tried</button>
                    <button class="mark-as-button" data-recipeid="${recipeID}" data-status="not-tried">Not Tried</button>
                    <button class="mark-as-button" data-recipeid="${recipeID}" data-status="to-cook">To Cook</button>
                    <button class="delete-button" data-recipeid="${recipeID}">Delete</button>
                `;

                recipeListContainer.appendChild(recipeDiv);
            });
        } else {
            recipeListContainer.innerHTML = '<p>No recipes found.</p>';
        }
    }
    
    const newForm = document.getElementById('new-form');

    newForm.addEventListener('submit', function (event) {
        event.preventDefault();

        // Create a FormData object from the form
        const formData = new FormData(newForm);

        // Make a POST request to the server
        fetch('/recipe', {
            method: 'POST',
            body: formData,
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                // Handle the response data, if needed
                console.log('Recipe added successfully:', data);

                window.location.href = '/recipes'; // Redirect to the recipe list page
            })
            .catch(error => console.error('Error:', error));
    });

    function handleFormSubmitWithFetch(url, method) {
        return function(event) {
            event.preventDefault();
            const formData = new FormData(event.target);
    
            fetch(url, { method: method, body: formData })
                .then(response => {
                    if (!response.ok) {
                        throw new Error(`HTTP error! status: ${response.status}`);
                    }
                    return response.text();
                })
                .then(data => {
                    console.log('Operation successful:', data);
                    window.location.href = '/recipes'; // Redirect or update the UI as needed
                })
                .catch(error => console.error('Error:', error));
        };
        
    }