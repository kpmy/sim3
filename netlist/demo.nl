import: [SM3, SM3r]
entities: 
  sr: SM3r
  p0: PROBE
  p1: PROBE
  s: SM3
  p2: PROBE
  p3: PROBE
netlist:
  T: [sr: A]
  N: [sr: B]
  spsr: [sr: Sr, p0: I]
  spcr: [sr: Cr, p1: I]
  T: [s: A]
  N: [s: B]
  sps: [s: S, p2: I]
  spc: [s: C, p3: I]