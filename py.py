from flask import Flask, request, redirect
# pip install GroupyAPI
from groupy import Client

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

        ans += f"""
        <a href="/groups?id={group.data['id']}&access_token={token}">{group.name} (click to create new bot)</a>
        <ol>
        """
        for bot in bots:
            if bot.data['group_id'] == group.data['id']:
                ans += f'''
                <li> <a href="/bots?id={group.data['id']}&access_token={token}">{bot.data['name']} (edit)</a></li>
                '''
        ans += '</ol>'


    print(bots[0].data)
    print(dir(bots[0]))

    return ans

@app.route("/groups")
def funccccction():
    token = request.args.get("access_token")
    group = request.args.get("id")
    c = Client.from_token(token)



app.run(port=8097,debug=True)
