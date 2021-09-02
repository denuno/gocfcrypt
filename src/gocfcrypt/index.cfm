<html>
<cfscript>
    plaintextFile = './plaintext.txt';
    keyFile = './keyfile';
    encryptedFile = './encrypted.txt';
    goEncryptedFile = './goencrypted.txt';
    goMessage = "";
    encryptedMessage = "";
    decryptedMessage = "";
    decryptedGoMessage = "";

    if(!fileExists(plaintextFile)){
        fileWrite(plaintextFile, "hello world");
    }

    if(!fileExists(keyFile)){
        fileWrite(keyFile, generateSecretKey('AES'));
        writeoutput("generating");
    }

    key = fileRead(keyFile);
    msg = fileRead(plaintextFile);

    if(!fileExists(encryptedFile)){
        encryptedMessage = encrypt( msg, key, 'AES/CBC/PKCS5Padding', 'HEX');
        filewrite(encryptedFile, encryptedMessage);
        writeoutput("enc");
    } else {
        writeoutput("reading go");
        encryptedMessage = fileread(encryptedFile);
    }


    if(!isNull(form.text)){
        decryptedMessage = decrypt( form.text, key, 'AES/CBC/PKCS5Padding', 'HEX');
        writeOutput(decryptedMessage & "<br/>");
    } else {
        decryptedMessage = decrypt( trim(encryptedMessage), key, 'AES/CBC/PKCS5Padding', 'HEX');
    }

    if(fileExists(goEncryptedFile)){
        goMessage = fileread(goEncryptedFile);
        try{
            decryptedGoMessage = decrypt( trim(goMessage), key, 'AES/CBC/PKCS5Padding', 'HEX');
        } catch(any e){
            decryptedGoMessage = e.message;
        }
    }
//    decMsg = decrypt( fileRead('./plaintext.aes'), key, 'AES/CBC/PKCS5Padding', 'HEX');

</cfscript>
<cfoutput>
<form action="?" method="post">
    <a href="?">reset</a>
    <pre>
        Message:#encryptedMessage#
        Decrypted:#decryptedMessage#

        GoMessage:#goMessage#
        GoDecryptedMessage:#decryptedGoMessage#
    </pre>
    <input type="text" name="text" value="#encryptedMessage#"/>
    <input type="submit">
</form>
</cfoutput>

</html>
