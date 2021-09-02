<html>
<cfscript>
    plaintext = './plaintext.txt';
    keyfile = './keyfile';
    encrypted = './encrypted.txt';
    goencrypted = './goencrypted.txt';

    if(!fileExists(plaintext)){
        fileWrite(plaintext, "hello world");
    }

    if(!fileExists(keyfile)){
        fileWrite(keyfile, generateSecretKey('AES'));
        writeoutput("generating");
    }

    if(!fileExists(encrypted)){
        encMsg = encrypt( msg, key, 'AES/CBC/PKCS5Padding', 'HEX');
        filewrite("./encrypted.txt", encmsg);
        writeoutput("enc");
    } else {
        writeoutput("reading go");
        encMsg = fileread(encrypted);
    }

    msg = fileRead(plaintext);
    key = fileRead(keyfile);

    if(!isNull(form.text)){
        decMsg = decrypt( form.text, key, 'AES/CBC/PKCS5Padding', 'HEX');
        writeOutput(decMsg & "<br/>");
    } else {
        decMsg = decrypt( trim(encMsg), key, 'AES/CBC/PKCS5Padding', 'HEX');
    }

    if(fileExists(goencrypted)){
        try{
            decGoMsg = decrypt( trim(fileread(goencrypted)), key, 'AES/CBC/PKCS5Padding', 'HEX');
        } catch(any e){
            decGoMsg = e.message;
        }
    }
//    decMsg = decrypt( fileRead('./plaintext.aes'), key, 'AES/CBC/PKCS5Padding', 'HEX');

</cfscript>
<form action="?" method="post">
    <a href="?">reset</a>
    <pre>Decrypted:<cfoutput>#decMsg#</cfoutput></pre>
    <pre>GoCrypted:<cfoutput>#decGoMsg#</cfoutput></pre>
    <input type="text" name="text" value="<cfoutput>#encMsg#</cfoutput>"/>
    <input type="submit">
</form>

</html>
