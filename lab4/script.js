let serials = [
    {
        "name": "Stranger things",
        "country": "U.S.A.",
        "episodes": 42
    },
    {
        "name": "The Vampire Diaries",
        "country": "U.S.A.",
        "episodes": 171
    },
    {
        "name": "Sen Çal Kapımı",
        "country": "Turkey",
        "episodes": 161
    }
]
let sum = 0;

for (let i = 0; i < serials.length; i++) {
    sum += serials[i]['episodes'];
}
avg_episodes = sum / serials.length;

function createSerial(name, country, episodes) {
    const serial = {
        "name": name,
        "country": country,
       "episodes": episodes,
        greeting: function () {
            return this.name + ' is my favourite and its from ' + this.country;
        }
    };
    return serial;
}

const Scream_queens = createSerial("Scream_queens", "U.S.A", 23);
console.log(Scream_queens.greeting());
serials.push(Scream_queens);

for (let i = 0; i < serials.length; i++) {
    document.write("<h5>Serial: " + serials[i]["name"] + "</h5>");
    document.write("<h5>Country: " + serials[i]["country"] + "</h5>");
    document.write("<h5>Episodes: " + serials[i]["episodes"] + "");
    document.writeln("");
}

document.write("<h5>Average duration: " + avg_episodes + " minutes</h5>");