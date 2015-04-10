library sim3.base;

import 'dart:html';

class Sim3Worker{

  Worker inner;

  Sim3Worker(){
    this.inner = new Worker("packages/sim3/src/sim3.js");
    this.inner.onMessage.listen((m){
      this.inner.postMessage({'Typ': 'init'});
    });
  }
}