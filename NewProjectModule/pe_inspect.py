import struct

def inspect(path):
    with open(path, 'rb') as f:
        data = f.read(512)
    e_lfanew = struct.unpack_from('<I', data, 0x3c)[0]
    pe_header = data[e_lfanew:e_lfanew+24]
    sig, machine, sections, timestamp, symptr, syms, optsize, characteristics = struct.unpack_from('<4sHHIIIHH', pe_header, 0)
    opt = data[e_lfanew+24:e_lfanew+24+optsize]
    magic, major, minor = struct.unpack_from('<HBB', opt, 0)
    subsystem = struct.unpack_from('<H', opt, 68)[0]
    print(path)
    print('sig', sig, 'machine', hex(machine), 'sections', sections, 'chars', hex(characteristics), 'magic', hex(magic), 'subsystem', hex(subsystem))

inspect('hello.exe')
inspect('newWindowApp.exe')
