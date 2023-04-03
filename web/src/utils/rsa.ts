// rsa
import JsEncrypt from 'jsencrypt';

export class RsaUtil {
  private publicKey: string;

  private privateKey: string;

  constructor() {
    // 公钥
    this.publicKey =
      'MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCi8+NXa27wtcVviKOCYVltuohvqRSwWXJ8gReSqH3C0avjCaoXGznpxCXcywiupLt/EuFYd6gqYQ0jRr/qVz5qZXnE3r4VO7pdLT3GcgCfFvES/1o6WQuCQC0sq1AbZj8vF2goKAmGlSahMDM4uqnTbIsBZ9XtP7x10asuyvnfYwIDAQAB';
    // 私钥
    this.privateKey = '';
  }

  encrypt(params: object | string): string | false {
    const Encrypt = new JsEncrypt();
    Encrypt.setPublicKey(this.publicKey);
    return Encrypt.encrypt(
      typeof params === 'object' ? JSON.stringify(params) : params
    );
  }

  decrypt(params: object | string) {
    const Decrypt = new JsEncrypt();
    Decrypt.setPrivateKey(this.privateKey);
    return Decrypt.decrypt(
      typeof params === 'object' ? JSON.stringify(params) : params
    );
  }
}

export default new RsaUtil();
