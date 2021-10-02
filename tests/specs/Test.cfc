component extends="testbox.system.BaseSpec" {

    plaintextFile = './plaintext.txt';
    keyFile = './keyfile';
    encryptedFile = './encrypted.txt';
    goEncryptedFile = './goencrypted.txt';

    function setup(){
        if(!fileExists(plaintextFile)){
            fileWrite(plaintextFile, "hello world");
        }
        if(!fileExists(keyFile)){
            fileWrite(keyFile, generateSecretKey('AES'));
        }
    }

    function testGoEncrypt(){
        key = fileRead(keyFile);
        msg = fileRead(plaintextFile);

        var encryptedMessage = encrypt( msg, key, 'AES/CBC/PKCS5Padding', 'HEX');
        var decryptedMessage = decrypt( encryptedMessage, key, 'AES/CBC/PKCS5Padding', 'HEX');

        expect( decryptedMessage ).toBe( msg );

    }

}