const express = require('express');

const app = express();
app.use(express.json());

app.post('/', (req, res) => {  
    const body = req.body;
    console.log('req: ', body);
    if (!body) {
        // bad request
        res.sendStatus(400);
        res.send("invalid");
        return;
    }

    const want = body["want_flag"];
    if (!want) {
        // bad request
        res.sendStatus(400);
        res.send("invalid");
        return;
    }

    // check
    const replaced = String(want).replace(/\!/g, '').replace(/\[/g, '').replace(/\]/g, '').replace(/\(/g, '').replace(/\)/g, '').replace(/\+/g, '');
    if (replaced.length > 0) {
        // forbidden
        console.log(replaced);
        res.sendStatus(403);
        res.send("invalid");
        return;
    }

    if (eval(want) === "yes"){
        res.send("taskctf{js_1s_4_tr1cky_l4ngu4ge}");
        return;
    }
    console.log(eval(want));
    res.send("invalid");
    return;
});

app.listen(30009, () => console.log('Listening on port 30009'));