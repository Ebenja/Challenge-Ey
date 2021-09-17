from flask import Flask, render_template, request, url_for, redirect, flash
from flask.wrappers import Response
import requests
import json
import pandas as pd
import matplotlib.pyplot as plt
from os import remove
from os import path

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
        
        params= {
            "dpi": int(dpi),
            "nombre": nombre,
            "apellidos": apellidos
        }
        
        response = requests.post(url+"User/create", json=params)
        print(response.status_code)
        if response.status_code == 200 or response.status_code == 201:
            flash('User created successfully')
            return redirect(url_for('addregister'))
        else:
            flash('User not created ')
            return redirect(url_for('addregister'))
            
        

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

@app.route('/delete/<id>', methods = ['POST','GET'])
def delete_user(id):
    # cur = mysql.connection.cursor()
    # cur.execute('DELETE FROM contacts WHERE id = {0}'.format(id))
    # mysql.connection.commit()
    response = requests.delete(url+"User/delete?dpi=" + id)
    print(url+"User/delete?dpi=" + id)
    if response.status_code >= 200 or response.status_code <= 200:
        flash('User Removed Successfully')
        return redirect(url_for('listUsers'))
    else:
        flash('User not Removed ')
        return redirect(url_for('listUsers'))

########### VACUNAS #########
@app.route("/forward/<id>/<vacuna>/<dosis>")
def move_forward(id, vacuna,dosis):
    params= {
            "dosis": int(id),
            "vacuna_id": int(vacuna),
            "persona_id": int(dosis)
        }
        
    response = requests.post(url+"Dosis/add", json=params)
    print(response.status_code)
    print(params)
    if response.status_code == 200 or response.status_code == 201:
        flash('Add Vaccine  successfully')
        return redirect(url_for('listUsers'))
    else:
        flash('Vaccine not added ')
        return redirect(url_for('listUsers'))

@app.route('/Vaccine/report')
def reportVaccine():
    # if request.method == 'GET':
    response2 = requests.get(url+"Dosis/report/")
    response = requests.get(url+"Dosis/All/")
    if response.status_code == 200 and response2.status_code == 200:
        users = response.json()
        reports = response2.json()
        df = pd.DataFrame(reports)
        df.plot(x="vacuna", kind="bar", y="dosis")
        plt.bar(df['vacuna'], df['dosis'])
        # fig = plt.figure(figsize=300)
        if path.exists('./static/plot1.png'):
            remove('./static/plot1.png')
        plt.savefig('./static/plot1.png')

        return render_template('VaccineReport.html',users=users, url ='/static/plot1.png' )
    else:
        return render_template('index.html')

if __name__ == "__main__":
    # global readEntrada
    app.run(debug=True)