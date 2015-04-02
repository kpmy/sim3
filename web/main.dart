// Copyright (c) 2015, <your name>. All rights reserved. Use of this source code
// is governed by a BSD-style license that can be found in the LICENSE file.

import 'dart:html';
import 'dart:async';
import 'dart:convert';

var Q = querySelector;

void main() {
  var get;
  get = (){
    HttpRequest.getString("http://localhost:3000/tri.json")
    ..then((s){
      var data = JSON.decode(s);
    })
    ..catchError((e){});
    new Future.delayed(new Duration(milliseconds: 100), get);
  };
  get();
}
