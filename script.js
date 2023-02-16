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
    sum += serials[i]['episodes']
}

avg_episodes = sum / serials.length;

let black_love = {
    "name": "Kara Sevda",
    "country": "Turkey",
    "episodes": 244
}


serials.push(black_love)

for (let i = 0; i < serials.length; i++) {
    document.write("<h5>Serial: " + serials[i]["name"] + "</h5>");
    document.write("<h5>Country: " + serials[i]["country"] + "</h5>");
    document.write("<h5>Episodes: " + serials[i]["episodes"] + "");
    document.writeln("")
}

document.write("<h5>Average duration: " + avg_episodes + " minutes</h5>")


