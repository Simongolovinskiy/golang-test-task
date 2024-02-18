const url = "https://api.coingecko.com/api/v3/coins/markets?" +
    "vs_currency=usd&order=market_cap_desc&per_page=250&page=1";

async function getData(url) {
    return fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .catch(error => {
            console.error('Fetch error:', error);
            throw error;
        });
}


async function displayTable() {
    const cryptoData = await getData(url);

    const tableBody = document.querySelector('#table tbody');

    cryptoData.forEach((currency, index) => {
        const row = document.createElement('tr');
        row.innerHTML = `<td>${currency.id}</td><td>${currency.symbol}</td><td>${currency.name}</td>`;

        if (currency.symbol === 'usdt') {
            row.classList.add('green-background');
        } else if (index <= 5) {
            row.classList.add('blue-background');
        }

        tableBody.appendChild(row);
    });
}

displayTable();
