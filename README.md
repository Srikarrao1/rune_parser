Runes are a form of data that can be embedded in Bitcoin transactions using the OP_RETURN script opcode. This opcode allows storing up to 80 bytes of arbitrary data in the blockchain. The stored data is immutable and can be easily retrieved from the blockchain.

Encoding Runes
Hexadecimal Encoding:
The data to be stored in a Bitcoin transaction’s OP_RETURN output is typically encoded in hexadecimal format.

Transaction Structure:
A Bitcoin transaction that includes Runes will have an OP_RETURN output.
The structure of this output is as follows:
6a: This is the hexadecimal representation of the OP_RETURN opcode.
Length of the data: The length of the data in bytes.
The actual Runes data: The data itself, encoded in hexadecimal format.


Here’s a step-by-step example:

Prepare the Data:

Convert your data to hexadecimal format.
Example data: "Hello, Runes!"
Hexadecimal representation: 48656c6c6f2c2052756e657321
Construct the OP_RETURN Output:

6a: OP_RETURN opcode in hexadecimal.
13: Length of the data in hexadecimal (19 bytes for "Hello, Runes!").
48656c6c6f2c2052756e657321: The actual data in hexadecimal.
The complete OP_RETURN output for this data would be:

6a1348656c6c6f2c2052756e657321

Decoding Runes
To decode a Runes transaction, follow these steps:
Extract the OP_RETURN Data:
Retrieve the transaction and identify the OP_RETURN output.
Decode the Hexadecimal Data:

Convert the hexadecimal data following the OP_RETURN opcode back to its original form.
For example, the hexadecimal string 48656c6c6f2c2052756e657321 converts back to "Hello, Runes!".
Example Workflow

Encoding:
Data: "Hello, Runes!"
Hexadecimal: 48656c6c6f2c2052756e657321
OP_RETURN Output: 6a1348656c6c6f2c2052756e657321


Decoding:
Extract the OP_RETURN data: 6a1348656c6c6f2c2052756e657321
Decode the hexadecimal part (48656c6c6f2c2052756e657321) back to the original data: "Hello, Runes!"
Retrieving Runes Data
To retrieve and decode Runes data from a Bitcoin transaction:

Use a blockchain explorer or a custom script to extract the OP_RETURN output.
Decode the hexadecimal data following the OP_RETURN opcode to its original form.
This process ensures that the Runes data is both immutable and easily accessible from the Bitcoin blockchain.
