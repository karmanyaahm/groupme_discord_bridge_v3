from flask import Flask, request, redirect

# pip install GroupyAPI
from groupy import Client

mycallback = "http://127.0.0.1:8097/"

app = Flask(__name__)


@app.route("/")
def hello_world():
    return redirect(
        "https://oauth.groupme.com/oauth/authorize?client_id=M9HtCkp8TW1602ecGtCE3xJjgz7aLCxyOXHCRtRR04ODGdnz",
        301,
    )


@app.route("/callback")
def callback():
    token = request.args.get("access_token")
    c = Client.from_token(token)
    groups = list(c.groups.list_all())
    bots = list(c.bots.list())
    ans = ""
    for group in groups:
        f = False
        for bot in bots:
            if (
                bot.data["group_id"] == group.data["id"]
                and bot.data["name"] == "Discord Sync"
            ):
                ans += f"{group.name} <ol>"
                f = True
        if not f:
            ans += f"""
            <a href="/groups?id={group.data['id']}&access_token={token}">{group.name} (click to create new bot)</a>
            <ol>
            """
        for bot in bots:
            if bot.data["group_id"] == group.data["id"]:
                ans += f"""
                <li> <a href="/bots?id={bot.data['bot_id']}&access_token={token}">{bot.data['name']} (edit)</a></li>
                """
        ans += "</ol>"

    return ans


@app.route("/groups")
def funccccction():
    token = request.args.get("access_token")
    group = request.args.get("id")
    c = Client.from_token(token)
    b = c.bots.create(
        "Discord Sync",
        group,
        avatar_url="https://discord.com/assets/2c21aeda16de354ba5334551a883b481.png",
        callback_url=mycallback,
        dm_notification=False,
    )
    return f"""
    groupme_bot_id = "{b.data['bot_id']}"<br>
    groupme_group_id = "{group}"
    <br><a href="/callback?access_token={token}">go back</a>
                """


@app.route("/bots")
def funcccccction():
    token = request.args.get("access_token")
    idd = request.args.get("id")
    delete = bool(request.args.get("delete", False))
    c = Client.from_token(token)

    for bot in c.bots.list():
        if idd == bot.data["bot_id"]:
            if delete:
                bot.destroy()
                return redirect(f"/callback?access_token={token}", 301)
            else:
                return f"""
                    name: {bot.data['name']}<br>
                    bot_id: {bot.data['bot_id']}<br>
                    bot_group: {bot.data['group_name']}<br>
                    callback_url: {bot.data['callback_url']}<br>
                    image: <img src="{bot.data['avatar_url']}"><br>
                <br><br><h3>configuration:</h3>
                    groupme_bot_id = "{bot.data['bot_id']}"<br>
                    groupme_group_id = "{bot.data['group_id']}"
                    <br><br><br><br><a href="/callback?access_token={token}">go back</a>
                    <br><a href="/bots?id={idd}&access_token={token}&delete=true">delete bot</a>
                    """


app.run(port=8097, debug=True)

