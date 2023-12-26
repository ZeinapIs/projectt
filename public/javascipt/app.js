// Создаем функцию для отправки GET-запроса на серверное API
function checkAPI() {
    fetch('/api/recipes')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json(); // Преобразуем ответ в JSON
        })
        .then(data => {
            // Проверяем данные на наличие свойств title и ID для каждого рецепта
            if (Array.isArray(data)) {
                for (const recipe of data) {
                    if (!recipe.hasOwnProperty('title') || !recipe.hasOwnProperty('ID')) {
                        console.error('API response does not contain expected properties.');
                        return;
                    }
                }
                console.log('API response is as expected.');
            } else {
                console.error('API response is not an array.');
            }
        })
        .catch(error => console.error('Error:', error));
}

// Вызываем функцию для проверки API при загрузке страницы
document.addEventListener('DOMContentLoaded', checkAPI);



// Другие функции и обработчики остаются неизменными.

document.addEventListener('DOMContentLoaded', function () 
    {
        fetch('/api/recipes')
            .then(response => response.json())
            .then(recipes => {
                const selectElement = document.getElementById('recipe-select');
                recipes.forEach(recipe => {
                    const option = new Option(recipe.title, recipe.ID);
                    selectElement.add(option);
                });
            })
            .catch(error => console.error('Error:', error));
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
    

    document.getElementById('delete-form').addEventListener('submit', function(event) {
        event.preventDefault();
        const recipeID = document.getElementById('recipe-id').value;
        deleteRecipe(recipeID);
    });
    

    if (newForm) {
        newForm.addEventListener('submit', handleNewRecipeSubmit);
    }
  
}
// Get deleteButton element and the current fact ID to delete
const deleteButton = document.querySelector('#delete-button')
const recipeToDelete = deleteButton && deleteButton.dataset.recipeid

// Add event listener to listen for button click
deleteButton && deleteButton.addEventListener('click', () => {
	// We ask the user if they are sure they want to delete the fact
    const result = confirm("Are you sure you want to delete this recipe?")
    
	// If the user cancels the prompt, we exit here
    if (!result) return

    // If the user confirms that they want to delete, we send a DELETE request
    // URL uses the current fact's ID
    // Lastly, we redirect to index
    return fetch(`/recipe/${recipeToDelete}`, { method: 'DELETE' })
            .then(() => document.location.href="/")
})




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
    
    
// Get the editForm element and the current recipe ID to edit
const editForm = document.querySelector('#form-update-recipe');
const recipeToEdit = editForm && editForm.dataset.recipeid;

// Add an event listener to listen for the form submit
editForm && editForm.addEventListener('submit', (event) => {
    // Prevent the default behaviour of the form element
    event.preventDefault();

    // Convert form data into a JavaScript object
    const formData = Object.fromEntries(new FormData(editForm));

    return fetch(`/recipe/${recipeToEdit}`, {
        // Use the PATCH or PUT method, depending on your API
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        // Convert the form's Object data into JSON
        body: JSON.stringify(formData),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        // Redirect to the updated recipe's detail view, or handle according to your app's logic
        document.location.href = `/recipe/${recipeToEdit}`;
    })
    .catch((error) => {
        console.error('There has been a problem with your fetch operation:', error);
        // Handle errors here, such as displaying a message to the user
    });
});

    

 

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
function handleRecipeListClick(event) {
    const target = event.target;
    if (target.classList.contains('delete-button')) {
        const recipeID = target.getAttribute('data-recipeid');
        deleteRecipe(recipeID);
    } else if (target.classList.contains('mark-as-button')) {
        const recipeID = target.getAttribute('data-recipeid');
        const status = target.getAttribute('data-status');
        updateRecipeStatus(recipeID, status);
    } else if (target.tagName === 'P') {
        const recipeID = target.getAttribute('data-recipeid');
        showRecipeDetails(recipeID);
    }
}

function handleDeleteRecipeSubmit(event) {
    event.preventDefault();
    const recipeID = document.getElementById('recipe-select').value;
    deleteRecipe(recipeID);
}

function handleSearchFormSubmit(event) {
    event.preventDefault();
    const query = document.getElementById('searchInput').value;
    const searchType = document.getElementById('searchType').value;
    redirectToSearchResults(query, searchType);
}

function updateRecipeList(recipes) {
    const recipeListContainer = document.getElementById('recipeList');
    recipeListContainer.innerHTML = '';

    if (recipes.length > 0) {
        recipes.forEach(recipe => {
            const recipeDiv = document.createElement('div');
            const recipeID = recipe.ID;

            recipeDiv.innerHTML = `
                <p data-recipeid="${recipeID}" style="cursor: pointer;">
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