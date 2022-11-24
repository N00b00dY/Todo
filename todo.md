# Todos

##
Check your todos and mark them as solved 
    - for each todo which is dynamicly generated, i need a checkbox
    - if I check the checkbox, a post is send which changedes the Activ status
        - if its 1 -> change to 0 
        - if its 0 -> change to 1
        - to work i have to send the ID of the checkbox so it migth be best solution
        to loop throut todos and create seperate eventlistener to the IDS. So its easy
        to send the ID
    - the post is send to distributon service via json api
    - from distribution service the Update function in DB service is called.