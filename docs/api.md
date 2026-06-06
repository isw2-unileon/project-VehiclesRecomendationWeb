# API Reference

The base URL locally is `http://localhost:8080` and in production `https://project-vehiclesrecomendationweb.onrender.com`. All communication between the Vue frontend (`public/index.html`) and the Go backend uses JSON over HTTP. The frontend is built with Vue 3 and makes all requests using the browser `fetch` API. Protected endpoints require the JWT token that is stored in `this.token` after login and sent in the `Authorization` header as `Bearer <token>`.

## Authentication

The login method validates the email format on the frontend before sending the request. It calls `POST /api/auth/login` with the user's email and password. On success, the server returns a JWT token and the username, which are stored in the Vue reactive state and in `localStorage` to survive page refreshes.

```javascript
const res = await fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email: this.logEmail, password: this.logPassword }),
});
const data = await res.json();
this.token = data.token;
this.currentUsername = data.username || this.logEmail.split('@')[0];
localStorage.setItem('token', this.token);
localStorage.setItem('username', this.currentUsername);
```

The register method applies two validations before sending anything to the server: it checks that the email matches a standard email pattern, and that the password is at least 8 characters long and includes at least one uppercase letter, one lowercase letter, one digit, and one special character. Only if both validations pass does it call `POST /api/auth/register`.

```javascript
const res = await fetch('/api/auth/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
        username: this.regUsername,
        email: this.regEmail,
        password: this.regPassword,
    }),
});
```

Logout does not call any endpoint. It clears the token and username from both the Vue state and `localStorage`, and resets all UI state — search results, AI response, favorites panel, form fields — back to their initial values.

## Database search

The database search calls `GET /api/cars/search` with up to four optional query parameters: `brand`, `fuel_type`, `max_price`, and `min_seats`. The frontend builds the query string dynamically using `URLSearchParams`, only appending a parameter if the user has filled in that field. The available fuel type options in the interface are Petrol, Diesel, and Hybrid.

```javascript
const params = new URLSearchParams();
if (this.searchBrand) params.append('brand', this.searchBrand);
if (this.searchFuel) params.append('fuel_type', this.searchFuel);
if (this.searchPrice) params.append('max_price', this.searchPrice);
if (this.searchSeats) params.append('min_seats', this.searchSeats);

const res = await fetch(`/api/cars/search?${params.toString()}`, { method: 'GET' });
this.searchResults = (await res.json()) || [];
```

Each car in the results shows its brand, model, fuel type, seats, and price. Three actions are available per result: opening a Google Images search for a photo of that car, viewing its full technical details, and saving it to favorites.

## Text search

The text search calls `GET /api/cars/text-search` with a single `q` parameter containing whatever the user typed. The input is encoded with `encodeURIComponent` to handle special characters. It searches across brand and model name fields and returns all matching vehicles with their engine, horsepower, fuel type, and seat count visible in the results.

```javascript
const res = await fetch(
    `/api/cars/text-search?q=${encodeURIComponent(this.textSearchInput)}`,
    { method: 'GET' }
);
this.textSearchResults = (await res.json()) || [];
```

## AI recommendation

The AI recommendation calls `POST /api/recommend` with the user's free-text prompt in the `preferences` field. This endpoint is protected and requires the JWT token. The server fetches relevant cars from the database, injects them into the prompt, and sends it to Groq's Llama 3.1 8B Instant model. The response contains a `recommendation` string which is displayed directly in the interface.

```javascript
const res = await fetch('/api/recommend', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + this.token,
    },
    body: JSON.stringify({ preferences: this.aiPrompt }),
});
const data = await res.json();
this.aiResponse = data.recommendation || 'System Error: ' + data.error;
```

## Favorites

The favorites panel loads automatically after login by calling `GET /api/favorites` with the JWT token. It reloads every time the user adds or removes a vehicle.

```javascript
const res = await fetch('/api/favorites', {
    method: 'GET',
    headers: { Authorization: 'Bearer ' + this.token },
});
this.favorites = (await res.json()) || [];
```

Adding a vehicle calls `POST /api/favorites` with the car's ID in the request body. The endpoint is protected.

```javascript
const res = await fetch('/api/favorites', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + this.token,
    },
    body: JSON.stringify({ car_id: carId }),
});
```

Removing a vehicle calls `DELETE /api/favorites` with the car ID as a query parameter. After a successful deletion the favorites list reloads automatically.

```javascript
const res = await fetch(`/api/favorites?car_id=${carId}`, {
    method: 'DELETE',
    headers: { Authorization: 'Bearer ' + this.token },
});
```