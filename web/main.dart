// Copyright (c) 2015, <your name>. All rights reserved. Use of this source code
// is governed by a BSD-style license that can be found in the LICENSE file.

import 'dart:html';
import 'dart:async';
import 'dart:convert';

var Q = querySelector;

class Item implements Comparable{
  int timestamp;
  String name;
  String type;
  bool meta;
  bool signal;

  @override
  int compareTo(Item that){
    DateTime a = new DateTime.fromMillisecondsSinceEpoch(this.timestamp);
    DateTime b = new DateTime.fromMillisecondsSinceEpoch(that.timestamp);
    if(a.isBefore(b))
      return -1;
    else if(a.isAfter(b))
      return 1;
    else return 0;
  }

  Item(this.timestamp, this.name, this.type, this.meta, this.signal);
}

List<Item> cache = new List();

update() async{
  (Q("#data") as DivElement).children.clear();
  Map<String, List<Item>> map = new Map();
  cache.forEach((i){
    if(!map.containsKey(i.name)){
      map[i.name] = new List<Item>();
    }
    map[i.name].add(i);
  });
  var keys = map.keys.toList();
  keys.sort();
  keys.forEach((s){
    var id = "p"+s.hashCode.toString();
    (Q("#data") as DivElement).appendHtml("<div>$s:</div><div id='p$id' style='white-space: nowrap;'></div>");
    String dump = "";
    map[s].forEach((i){
      var sig = i.signal == null ? "0" : i.signal ? "+" : "-";
      if(i.meta)
        dump = "$sig $dump";
      else
        dump = "_  $dump";
    });
    (Q("#p$id") as DivElement).appendHtml(dump);
  });
  return 0;
}

main() async{
  var get;
  get = (){
    HttpRequest.getString("http://localhost:3000/tri.json")
    ..then((s){
      var data = JSON.decode(s);
      List il = (data as List);
      il.forEach((i){
        bool sig = i["Signal"]["N"] ? null : i["Signal"]["T"];
        cache.add(new Item(i["Timestamp"], i["Name"], i["Type"], i["Meta"], sig));
      });
      cache.sort();
      update();
    })
    ..catchError((e){});
    new Future.delayed(new Duration(milliseconds: 300), get);
  };
  get();
}
