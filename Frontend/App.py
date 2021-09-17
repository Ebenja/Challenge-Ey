from flask import Flask, render_template, request, url_for, redirect, flash
from flask.wrappers import Response
import requests
import json
url = "http://localhost:3000/"
app = Flask(__name__)

app.secret_key = "mysecretkey"

@app.route("/")
def home():
    return render_template('index.html')

@app.route('/User/Create')
def addregister():
    return render_template('Create.html')

@app.route('/CreateDone', methods=['POST'])
def add_contact():
    if request.method == 'POST':
        dpi = request.form['dpi']
        nombre = request.form['nombre']
        apellidos = request.form['apellidos']
        # cur = mysql.connection.cursor()
        # cur.execute("INSERT INTO contacts (fullname, phone, email) VALUES (%s,%s,%s)", (fullname, phone, email))
        # mysql.connection.commit()
        # flash('Contact Added successfully')
        params= {
            "dpi": int(dpi),
            "nombre": nombre,
            "apellidos": apellidos
        }
        print(params)
        response = requests.post(url+"User/create", json=params)
        print(response.status_code)
        if response.status_code == 200 or response.status_code == 201:
            flash('User created successfully')
            return redirect(url_for('addregister'))
        else:
            flash('User not created ')
            return redirect(url_for('addregister'))
        # return redirect(url_for('Index'))
        

@app.route('/User/list')
def listUsers():
    # if request.method == 'GET':
    response = requests.get(url+"User/getAll/")
    if response.status_code == 200:
        users = response.json()
        print(response.json())
        return render_template('ListUser.html',users=users )
    else:
        return render_template('index.html')
        
@app.route('/User/edit/<id>', methods = ['POST', 'GET'])
def get_contact(id):
    response = requests.get(url+"User/get?dpiuser=" + id)
    if response.status_code == 200:
        users = response.json()
        print("usuario 0:")
        # print(users[0])
        print(response.json())
        return render_template('Edit.html',user=users )
    else:
        return render_template('index.html')


if __name__ == "__main__":
    # global readEntrada
    app.run(debug=True)