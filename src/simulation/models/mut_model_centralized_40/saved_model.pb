ев
°╩
,
Abs
x"T
y"T"
Ttype:

2	
:
Add
x"T
y"T
z"T"
Ttype:
2	
x
Assign
ref"TА

value"T

output_ref"TА"	
Ttype"
validate_shapebool("
use_lockingbool(Ш
s
	AssignAdd
ref"TА

value"T

output_ref"TА" 
Ttype:
2	"
use_lockingbool( 
R
BroadcastGradientArgs
s0"T
s1"T
r0"T
r1"T"
Ttype0:
2	
N
Cast	
x"SrcT	
y"DstT"
SrcTtype"
DstTtype"
Truncatebool( 
8
Const
output"dtype"
valuetensor"
dtypetype
S
DynamicStitch
indices*N
data"T*N
merged"T"
Nint(0"	
Ttype
^
Fill
dims"
index_type

value"T
output"T"	
Ttype"

index_typetype0:
2	
?
FloorDiv
x"T
y"T
z"T"
Ttype:
2	
9
FloorMod
x"T
y"T
z"T"
Ttype:

2	
.
Identity

input"T
output"T"	
Ttype
:
InvertPermutation
x"T
y"T"
Ttype0:
2	
N
IsVariableInitialized
ref"dtypeА
is_initialized
"
dtypetypeШ
q
MatMul
a"T
b"T
product"T"
transpose_abool( "
transpose_bbool( "
Ttype:

2	
;
Maximum
x"T
y"T
z"T"
Ttype:

2	Р
Н
Mean

input"T
reduction_indices"Tidx
output"T"
	keep_dimsbool( " 
Ttype:
2	"
Tidxtype0:
2	
e
MergeV2Checkpoints
checkpoint_prefixes
destination_prefix"
delete_old_dirsbool(И
;
Minimum
x"T
y"T
z"T"
Ttype:

2	Р
=
Mul
x"T
y"T
z"T"
Ttype:
2	Р
.
Neg
x"T
y"T"
Ttype:

2	

NoOp
E
NotEqual
x"T
y"T
z
"
Ttype:
2	
Р
M
Pack
values"T*N
output"T"
Nint(0"	
Ttype"
axisint 
C
Placeholder
output"dtype"
dtypetype"
shapeshape:
X
PlaceholderWithDefault
input"dtype
output"dtype"
dtypetype"
shapeshape
6
Pow
x"T
y"T
z"T"
Ttype:

2	
Н
Prod

input"T
reduction_indices"Tidx
output"T"
	keep_dimsbool( " 
Ttype:
2	"
Tidxtype0:
2	
Е
RandomStandardNormal

shape"T
output"dtype"
seedint "
seed2int "
dtypetype:
2"
Ttype:
2	И
a
Range
start"Tidx
limit"Tidx
delta"Tidx
output"Tidx"
Tidxtype0:	
2	
>
RealDiv
x"T
y"T
z"T"
Ttype:
2	
E
Relu
features"T
activations"T"
Ttype:
2	
V
ReluGrad
	gradients"T
features"T
	backprops"T"
Ttype:
2	
[
Reshape
tensor"T
shape"Tshape
output"T"	
Ttype"
Tshapetype0:
2	
o
	RestoreV2

prefix
tensor_names
shape_and_slices
tensors2dtypes"
dtypes
list(type)(0И
l
SaveV2

prefix
tensor_names
shape_and_slices
tensors2dtypes"
dtypes
list(type)(0И
P
Shape

input"T
output"out_type"	
Ttype"
out_typetype0:
2	
H
ShardedFilename
basename	
shard

num_shards
filename
/
Sign
x"T
y"T"
Ttype:

2	
-
Sqrt
x"T
y"T"
Ttype:

2
1
Square
x"T
y"T"
Ttype:

2	
N

StringJoin
inputs*N

output"
Nint(0"
	separatorstring 
:
Sub
x"T
y"T
z"T"
Ttype:
2	
М
Sum

input"T
reduction_indices"Tidx
output"T"
	keep_dimsbool( " 
Ttype:
2	"
Tidxtype0:
2	
c
Tile

input"T
	multiples"
Tmultiples
output"T"	
Ttype"

Tmultiplestype0:
2	
P
	Transpose
x"T
perm"Tperm
y"T"	
Ttype"
Tpermtype0:
2	
P
Unpack

value"T
output"T*num"
numint("	
Ttype"
axisint 
s

VariableV2
ref"dtypeА"
shapeshape"
dtypetype"
	containerstring "
shared_namestring И"mut_tag*1.13.12b'v1.13.1-0-g6612da8951'ък

\
keras_learning_phase/inputConst*
value	B
 Z *
dtype0
*
_output_shapes
: 
|
keras_learning_phasePlaceholderWithDefaultkeras_learning_phase/input*
_output_shapes
: *
shape: *
dtype0

|
input_layer_inputPlaceholder*+
_output_shapes
:         z* 
shape:         z*
dtype0
p
input_layer/random_normal/shapeConst*
valueB"z      *
dtype0*
_output_shapes
:
c
input_layer/random_normal/meanConst*
valueB
 *    *
dtype0*
_output_shapes
: 
e
 input_layer/random_normal/stddevConst*
valueB
 *═╠L=*
dtype0*
_output_shapes
: 
╜
.input_layer/random_normal/RandomStandardNormalRandomStandardNormalinput_layer/random_normal/shape*
T0*
dtype0*
_output_shapes
:	zА*
seed2злп*
seed▒ х)
а
input_layer/random_normal/mulMul.input_layer/random_normal/RandomStandardNormal input_layer/random_normal/stddev*
T0*
_output_shapes
:	zА
Й
input_layer/random_normalAddinput_layer/random_normal/mulinput_layer/random_normal/mean*
T0*
_output_shapes
:	zА
И
input_layer/kernel
VariableV2*
_output_shapes
:	zА*
	container *
shape:	zА*
shared_name *
dtype0
╠
input_layer/kernel/AssignAssigninput_layer/kernelinput_layer/random_normal*
use_locking(*
T0*%
_class
loc:@input_layer/kernel*
validate_shape(*
_output_shapes
:	zА
И
input_layer/kernel/readIdentityinput_layer/kernel*
T0*%
_class
loc:@input_layer/kernel*
_output_shapes
:	zА
`
input_layer/ConstConst*
valueBА*    *
dtype0*
_output_shapes	
:А
~
input_layer/bias
VariableV2*
_output_shapes	
:А*
	container *
shape:А*
shared_name *
dtype0
║
input_layer/bias/AssignAssigninput_layer/biasinput_layer/Const*#
_class
loc:@input_layer/bias*
validate_shape(*
_output_shapes	
:А*
use_locking(*
T0
~
input_layer/bias/readIdentityinput_layer/bias*
_output_shapes	
:А*
T0*#
_class
loc:@input_layer/bias
b
input_layer/ShapeShapeinput_layer_input*
T0*
out_type0*
_output_shapes
:
l
input_layer/unstackUnpackinput_layer/Shape*
_output_shapes
: : : *	
num*
T0*

axis 
d
input_layer/Shape_1Const*
valueB"z      *
dtype0*
_output_shapes
:
n
input_layer/unstack_1Unpackinput_layer/Shape_1*

axis *
_output_shapes
: : *	
num*
T0
j
input_layer/Reshape/shapeConst*
_output_shapes
:*
valueB"    z   *
dtype0
М
input_layer/ReshapeReshapeinput_layer_inputinput_layer/Reshape/shape*
T0*
Tshape0*'
_output_shapes
:         z
k
input_layer/transpose/permConst*
valueB"       *
dtype0*
_output_shapes
:
О
input_layer/transpose	Transposeinput_layer/kernel/readinput_layer/transpose/perm*
T0*
_output_shapes
:	zА*
Tperm0
l
input_layer/Reshape_1/shapeConst*
valueB"z       *
dtype0*
_output_shapes
:
М
input_layer/Reshape_1Reshapeinput_layer/transposeinput_layer/Reshape_1/shape*
T0*
Tshape0*
_output_shapes
:	zА
б
input_layer/MatMulMatMulinput_layer/Reshapeinput_layer/Reshape_1*
T0*(
_output_shapes
:         А*
transpose_a( *
transpose_b( 
_
input_layer/Reshape_2/shape/1Const*
value	B :*
dtype0*
_output_shapes
: 
`
input_layer/Reshape_2/shape/2Const*
value
B :А*
dtype0*
_output_shapes
: 
░
input_layer/Reshape_2/shapePackinput_layer/unstackinput_layer/Reshape_2/shape/1input_layer/Reshape_2/shape/2*
T0*

axis *
N*
_output_shapes
:
Ц
input_layer/Reshape_2Reshapeinput_layer/MatMulinput_layer/Reshape_2/shape*
T0*
Tshape0*,
_output_shapes
:         А
p
input_layer/Reshape_3/shapeConst*!
valueB"         *
dtype0*
_output_shapes
:
Р
input_layer/Reshape_3Reshapeinput_layer/bias/readinput_layer/Reshape_3/shape*
T0*
Tshape0*#
_output_shapes
:А
{
input_layer/addAddinput_layer/Reshape_2input_layer/Reshape_3*
T0*,
_output_shapes
:         А
`
input_layer/ReluReluinput_layer/add*
T0*,
_output_shapes
:         А
l
dense_1/random_normal/shapeConst*
valueB"   А   *
dtype0*
_output_shapes
:
_
dense_1/random_normal/meanConst*
dtype0*
_output_shapes
: *
valueB
 *    
a
dense_1/random_normal/stddevConst*
valueB
 *═╠L=*
dtype0*
_output_shapes
: 
╢
*dense_1/random_normal/RandomStandardNormalRandomStandardNormaldense_1/random_normal/shape*
T0*
dtype0* 
_output_shapes
:
АА*
seed2╕╜╥*
seed▒ х)
Х
dense_1/random_normal/mulMul*dense_1/random_normal/RandomStandardNormaldense_1/random_normal/stddev* 
_output_shapes
:
АА*
T0
~
dense_1/random_normalAdddense_1/random_normal/muldense_1/random_normal/mean* 
_output_shapes
:
АА*
T0
Ж
dense_1/kernel
VariableV2*
shared_name *
dtype0* 
_output_shapes
:
АА*
	container *
shape:
АА
╜
dense_1/kernel/AssignAssigndense_1/kerneldense_1/random_normal*
use_locking(*
T0*!
_class
loc:@dense_1/kernel*
validate_shape(* 
_output_shapes
:
АА
}
dense_1/kernel/readIdentitydense_1/kernel* 
_output_shapes
:
АА*
T0*!
_class
loc:@dense_1/kernel
\
dense_1/ConstConst*
valueBА*    *
dtype0*
_output_shapes	
:А
z
dense_1/bias
VariableV2*
dtype0*
_output_shapes	
:А*
	container *
shape:А*
shared_name 
к
dense_1/bias/AssignAssigndense_1/biasdense_1/Const*
_class
loc:@dense_1/bias*
validate_shape(*
_output_shapes	
:А*
use_locking(*
T0
r
dense_1/bias/readIdentitydense_1/bias*
T0*
_class
loc:@dense_1/bias*
_output_shapes	
:А
]
dense_1/ShapeShapeinput_layer/Relu*
T0*
out_type0*
_output_shapes
:
d
dense_1/unstackUnpackdense_1/Shape*
_output_shapes
: : : *	
num*
T0*

axis 
`
dense_1/Shape_1Const*
valueB"   А   *
dtype0*
_output_shapes
:
f
dense_1/unstack_1Unpackdense_1/Shape_1*	
num*
T0*

axis *
_output_shapes
: : 
f
dense_1/Reshape/shapeConst*
valueB"       *
dtype0*
_output_shapes
:
Д
dense_1/ReshapeReshapeinput_layer/Reludense_1/Reshape/shape*
T0*
Tshape0*(
_output_shapes
:         А
g
dense_1/transpose/permConst*
valueB"       *
dtype0*
_output_shapes
:
Г
dense_1/transpose	Transposedense_1/kernel/readdense_1/transpose/perm*
T0* 
_output_shapes
:
АА*
Tperm0
h
dense_1/Reshape_1/shapeConst*
_output_shapes
:*
valueB"       *
dtype0
Б
dense_1/Reshape_1Reshapedense_1/transposedense_1/Reshape_1/shape*
Tshape0* 
_output_shapes
:
АА*
T0
Х
dense_1/MatMulMatMuldense_1/Reshapedense_1/Reshape_1*
transpose_b( *
T0*(
_output_shapes
:         А*
transpose_a( 
[
dense_1/Reshape_2/shape/1Const*
dtype0*
_output_shapes
: *
value	B :
\
dense_1/Reshape_2/shape/2Const*
dtype0*
_output_shapes
: *
value
B :А
а
dense_1/Reshape_2/shapePackdense_1/unstackdense_1/Reshape_2/shape/1dense_1/Reshape_2/shape/2*
T0*

axis *
N*
_output_shapes
:
К
dense_1/Reshape_2Reshapedense_1/MatMuldense_1/Reshape_2/shape*
T0*
Tshape0*,
_output_shapes
:         А
l
dense_1/Reshape_3/shapeConst*!
valueB"      А   *
dtype0*
_output_shapes
:
Д
dense_1/Reshape_3Reshapedense_1/bias/readdense_1/Reshape_3/shape*#
_output_shapes
:А*
T0*
Tshape0
o
dense_1/addAdddense_1/Reshape_2dense_1/Reshape_3*
T0*,
_output_shapes
:         А
X
dense_1/ReluReludense_1/add*,
_output_shapes
:         А*
T0
l
dense_2/random_normal/shapeConst*
valueB"А   А   *
dtype0*
_output_shapes
:
_
dense_2/random_normal/meanConst*
valueB
 *    *
dtype0*
_output_shapes
: 
a
dense_2/random_normal/stddevConst*
valueB
 *═╠L=*
dtype0*
_output_shapes
: 
╡
*dense_2/random_normal/RandomStandardNormalRandomStandardNormaldense_2/random_normal/shape*
dtype0* 
_output_shapes
:
АА*
seed2░ь*
seed▒ х)*
T0
Х
dense_2/random_normal/mulMul*dense_2/random_normal/RandomStandardNormaldense_2/random_normal/stddev*
T0* 
_output_shapes
:
АА
~
dense_2/random_normalAdddense_2/random_normal/muldense_2/random_normal/mean*
T0* 
_output_shapes
:
АА
Ж
dense_2/kernel
VariableV2* 
_output_shapes
:
АА*
	container *
shape:
АА*
shared_name *
dtype0
╜
dense_2/kernel/AssignAssigndense_2/kerneldense_2/random_normal*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0*!
_class
loc:@dense_2/kernel
}
dense_2/kernel/readIdentitydense_2/kernel*
T0*!
_class
loc:@dense_2/kernel* 
_output_shapes
:
АА
\
dense_2/ConstConst*
_output_shapes	
:А*
valueBА*    *
dtype0
z
dense_2/bias
VariableV2*
shared_name *
dtype0*
_output_shapes	
:А*
	container *
shape:А
к
dense_2/bias/AssignAssigndense_2/biasdense_2/Const*
_class
loc:@dense_2/bias*
validate_shape(*
_output_shapes	
:А*
use_locking(*
T0
r
dense_2/bias/readIdentitydense_2/bias*
T0*
_class
loc:@dense_2/bias*
_output_shapes	
:А
Y
dense_2/ShapeShapedense_1/Relu*
T0*
out_type0*
_output_shapes
:
d
dense_2/unstackUnpackdense_2/Shape*
_output_shapes
: : : *	
num*
T0*

axis 
`
dense_2/Shape_1Const*
valueB"А   А   *
dtype0*
_output_shapes
:
f
dense_2/unstack_1Unpackdense_2/Shape_1*
_output_shapes
: : *	
num*
T0*

axis 
f
dense_2/Reshape/shapeConst*
valueB"    А   *
dtype0*
_output_shapes
:
А
dense_2/ReshapeReshapedense_1/Reludense_2/Reshape/shape*
T0*
Tshape0*(
_output_shapes
:         А
g
dense_2/transpose/permConst*
valueB"       *
dtype0*
_output_shapes
:
Г
dense_2/transpose	Transposedense_2/kernel/readdense_2/transpose/perm*
T0* 
_output_shapes
:
АА*
Tperm0
h
dense_2/Reshape_1/shapeConst*
_output_shapes
:*
valueB"А       *
dtype0
Б
dense_2/Reshape_1Reshapedense_2/transposedense_2/Reshape_1/shape*
T0*
Tshape0* 
_output_shapes
:
АА
Х
dense_2/MatMulMatMuldense_2/Reshapedense_2/Reshape_1*(
_output_shapes
:         А*
transpose_a( *
transpose_b( *
T0
[
dense_2/Reshape_2/shape/1Const*
value	B :*
dtype0*
_output_shapes
: 
\
dense_2/Reshape_2/shape/2Const*
value
B :А*
dtype0*
_output_shapes
: 
а
dense_2/Reshape_2/shapePackdense_2/unstackdense_2/Reshape_2/shape/1dense_2/Reshape_2/shape/2*
N*
_output_shapes
:*
T0*

axis 
К
dense_2/Reshape_2Reshapedense_2/MatMuldense_2/Reshape_2/shape*
Tshape0*,
_output_shapes
:         А*
T0
l
dense_2/Reshape_3/shapeConst*!
valueB"      А   *
dtype0*
_output_shapes
:
Д
dense_2/Reshape_3Reshapedense_2/bias/readdense_2/Reshape_3/shape*
T0*
Tshape0*#
_output_shapes
:А
o
dense_2/addAdddense_2/Reshape_2dense_2/Reshape_3*
T0*,
_output_shapes
:         А
X
dense_2/ReluReludense_2/add*,
_output_shapes
:         А*
T0
l
dense_3/random_normal/shapeConst*
dtype0*
_output_shapes
:*
valueB"А   А   
_
dense_3/random_normal/meanConst*
_output_shapes
: *
valueB
 *    *
dtype0
a
dense_3/random_normal/stddevConst*
valueB
 *═╠L=*
dtype0*
_output_shapes
: 
╢
*dense_3/random_normal/RandomStandardNormalRandomStandardNormaldense_3/random_normal/shape*
seed▒ х)*
T0*
dtype0* 
_output_shapes
:
АА*
seed2№м╒
Х
dense_3/random_normal/mulMul*dense_3/random_normal/RandomStandardNormaldense_3/random_normal/stddev*
T0* 
_output_shapes
:
АА
~
dense_3/random_normalAdddense_3/random_normal/muldense_3/random_normal/mean* 
_output_shapes
:
АА*
T0
Ж
dense_3/kernel
VariableV2*
shape:
АА*
shared_name *
dtype0* 
_output_shapes
:
АА*
	container 
╜
dense_3/kernel/AssignAssigndense_3/kerneldense_3/random_normal*
use_locking(*
T0*!
_class
loc:@dense_3/kernel*
validate_shape(* 
_output_shapes
:
АА
}
dense_3/kernel/readIdentitydense_3/kernel*!
_class
loc:@dense_3/kernel* 
_output_shapes
:
АА*
T0
\
dense_3/ConstConst*
valueBА*    *
dtype0*
_output_shapes	
:А
z
dense_3/bias
VariableV2*
shape:А*
shared_name *
dtype0*
_output_shapes	
:А*
	container 
к
dense_3/bias/AssignAssigndense_3/biasdense_3/Const*
use_locking(*
T0*
_class
loc:@dense_3/bias*
validate_shape(*
_output_shapes	
:А
r
dense_3/bias/readIdentitydense_3/bias*
T0*
_class
loc:@dense_3/bias*
_output_shapes	
:А
Y
dense_3/ShapeShapedense_2/Relu*
out_type0*
_output_shapes
:*
T0
d
dense_3/unstackUnpackdense_3/Shape*	
num*
T0*

axis *
_output_shapes
: : : 
`
dense_3/Shape_1Const*
_output_shapes
:*
valueB"А   А   *
dtype0
f
dense_3/unstack_1Unpackdense_3/Shape_1*	
num*
T0*

axis *
_output_shapes
: : 
f
dense_3/Reshape/shapeConst*
valueB"    А   *
dtype0*
_output_shapes
:
А
dense_3/ReshapeReshapedense_2/Reludense_3/Reshape/shape*
T0*
Tshape0*(
_output_shapes
:         А
g
dense_3/transpose/permConst*
valueB"       *
dtype0*
_output_shapes
:
Г
dense_3/transpose	Transposedense_3/kernel/readdense_3/transpose/perm* 
_output_shapes
:
АА*
Tperm0*
T0
h
dense_3/Reshape_1/shapeConst*
dtype0*
_output_shapes
:*
valueB"А       
Б
dense_3/Reshape_1Reshapedense_3/transposedense_3/Reshape_1/shape*
T0*
Tshape0* 
_output_shapes
:
АА
Х
dense_3/MatMulMatMuldense_3/Reshapedense_3/Reshape_1*(
_output_shapes
:         А*
transpose_a( *
transpose_b( *
T0
[
dense_3/Reshape_2/shape/1Const*
value	B :*
dtype0*
_output_shapes
: 
\
dense_3/Reshape_2/shape/2Const*
value
B :А*
dtype0*
_output_shapes
: 
а
dense_3/Reshape_2/shapePackdense_3/unstackdense_3/Reshape_2/shape/1dense_3/Reshape_2/shape/2*
_output_shapes
:*
T0*

axis *
N
К
dense_3/Reshape_2Reshapedense_3/MatMuldense_3/Reshape_2/shape*,
_output_shapes
:         А*
T0*
Tshape0
l
dense_3/Reshape_3/shapeConst*!
valueB"      А   *
dtype0*
_output_shapes
:
Д
dense_3/Reshape_3Reshapedense_3/bias/readdense_3/Reshape_3/shape*
T0*
Tshape0*#
_output_shapes
:А
o
dense_3/addAdddense_3/Reshape_2dense_3/Reshape_3*
T0*,
_output_shapes
:         А
X
dense_3/ReluReludense_3/add*,
_output_shapes
:         А*
T0
q
 output_layer/random_normal/shapeConst*
valueB"А   (   *
dtype0*
_output_shapes
:
d
output_layer/random_normal/meanConst*
_output_shapes
: *
valueB
 *    *
dtype0
f
!output_layer/random_normal/stddevConst*
valueB
 *═╠L=*
dtype0*
_output_shapes
: 
╛
/output_layer/random_normal/RandomStandardNormalRandomStandardNormal output_layer/random_normal/shape*
dtype0*
_output_shapes
:	А(*
seed2╢├)*
seed▒ х)*
T0
г
output_layer/random_normal/mulMul/output_layer/random_normal/RandomStandardNormal!output_layer/random_normal/stddev*
_output_shapes
:	А(*
T0
М
output_layer/random_normalAddoutput_layer/random_normal/muloutput_layer/random_normal/mean*
T0*
_output_shapes
:	А(
Й
output_layer/kernel
VariableV2*
shared_name *
dtype0*
_output_shapes
:	А(*
	container *
shape:	А(
╨
output_layer/kernel/AssignAssignoutput_layer/kerneloutput_layer/random_normal*
T0*&
_class
loc:@output_layer/kernel*
validate_shape(*
_output_shapes
:	А(*
use_locking(
Л
output_layer/kernel/readIdentityoutput_layer/kernel*
T0*&
_class
loc:@output_layer/kernel*
_output_shapes
:	А(
_
output_layer/ConstConst*
valueB(*    *
dtype0*
_output_shapes
:(
}
output_layer/bias
VariableV2*
dtype0*
_output_shapes
:(*
	container *
shape:(*
shared_name 
╜
output_layer/bias/AssignAssignoutput_layer/biasoutput_layer/Const*
use_locking(*
T0*$
_class
loc:@output_layer/bias*
validate_shape(*
_output_shapes
:(
А
output_layer/bias/readIdentityoutput_layer/bias*$
_class
loc:@output_layer/bias*
_output_shapes
:(*
T0
^
output_layer/ShapeShapedense_3/Relu*
T0*
out_type0*
_output_shapes
:
n
output_layer/unstackUnpackoutput_layer/Shape*
_output_shapes
: : : *	
num*
T0*

axis 
e
output_layer/Shape_1Const*
_output_shapes
:*
valueB"А   (   *
dtype0
p
output_layer/unstack_1Unpackoutput_layer/Shape_1*	
num*
T0*

axis *
_output_shapes
: : 
k
output_layer/Reshape/shapeConst*
valueB"    А   *
dtype0*
_output_shapes
:
К
output_layer/ReshapeReshapedense_3/Reluoutput_layer/Reshape/shape*(
_output_shapes
:         А*
T0*
Tshape0
l
output_layer/transpose/permConst*
dtype0*
_output_shapes
:*
valueB"       
С
output_layer/transpose	Transposeoutput_layer/kernel/readoutput_layer/transpose/perm*
_output_shapes
:	А(*
Tperm0*
T0
m
output_layer/Reshape_1/shapeConst*
valueB"А       *
dtype0*
_output_shapes
:
П
output_layer/Reshape_1Reshapeoutput_layer/transposeoutput_layer/Reshape_1/shape*
T0*
Tshape0*
_output_shapes
:	А(
г
output_layer/MatMulMatMuloutput_layer/Reshapeoutput_layer/Reshape_1*
transpose_b( *
T0*'
_output_shapes
:         (*
transpose_a( 
`
output_layer/Reshape_2/shape/1Const*
value	B :*
dtype0*
_output_shapes
: 
`
output_layer/Reshape_2/shape/2Const*
value	B :(*
dtype0*
_output_shapes
: 
┤
output_layer/Reshape_2/shapePackoutput_layer/unstackoutput_layer/Reshape_2/shape/1output_layer/Reshape_2/shape/2*
T0*

axis *
N*
_output_shapes
:
Ш
output_layer/Reshape_2Reshapeoutput_layer/MatMuloutput_layer/Reshape_2/shape*
T0*
Tshape0*+
_output_shapes
:         (
q
output_layer/Reshape_3/shapeConst*!
valueB"      (   *
dtype0*
_output_shapes
:
Т
output_layer/Reshape_3Reshapeoutput_layer/bias/readoutput_layer/Reshape_3/shape*
T0*
Tshape0*"
_output_shapes
:(
}
output_layer/addAddoutput_layer/Reshape_2output_layer/Reshape_3*
T0*+
_output_shapes
:         (
_
Adam/iterations/initial_valueConst*
_output_shapes
: *
value	B	 R *
dtype0	
s
Adam/iterations
VariableV2*
_output_shapes
: *
	container *
shape: *
shared_name *
dtype0	
╛
Adam/iterations/AssignAssignAdam/iterationsAdam/iterations/initial_value*
use_locking(*
T0	*"
_class
loc:@Adam/iterations*
validate_shape(*
_output_shapes
: 
v
Adam/iterations/readIdentityAdam/iterations*
T0	*"
_class
loc:@Adam/iterations*
_output_shapes
: 
Z
Adam/lr/initial_valueConst*
dtype0*
_output_shapes
: *
valueB
 *╖╤8
k
Adam/lr
VariableV2*
shared_name *
dtype0*
_output_shapes
: *
	container *
shape: 
Ю
Adam/lr/AssignAssignAdam/lrAdam/lr/initial_value*
use_locking(*
T0*
_class
loc:@Adam/lr*
validate_shape(*
_output_shapes
: 
^
Adam/lr/readIdentityAdam/lr*
T0*
_class
loc:@Adam/lr*
_output_shapes
: 
^
Adam/beta_1/initial_valueConst*
valueB
 *fff?*
dtype0*
_output_shapes
: 
o
Adam/beta_1
VariableV2*
dtype0*
_output_shapes
: *
	container *
shape: *
shared_name 
о
Adam/beta_1/AssignAssignAdam/beta_1Adam/beta_1/initial_value*
use_locking(*
T0*
_class
loc:@Adam/beta_1*
validate_shape(*
_output_shapes
: 
j
Adam/beta_1/readIdentityAdam/beta_1*
_output_shapes
: *
T0*
_class
loc:@Adam/beta_1
^
Adam/beta_2/initial_valueConst*
valueB
 *w╛?*
dtype0*
_output_shapes
: 
o
Adam/beta_2
VariableV2*
shape: *
shared_name *
dtype0*
_output_shapes
: *
	container 
о
Adam/beta_2/AssignAssignAdam/beta_2Adam/beta_2/initial_value*
T0*
_class
loc:@Adam/beta_2*
validate_shape(*
_output_shapes
: *
use_locking(
j
Adam/beta_2/readIdentityAdam/beta_2*
T0*
_class
loc:@Adam/beta_2*
_output_shapes
: 
]
Adam/decay/initial_valueConst*
dtype0*
_output_shapes
: *
valueB
 *    
n

Adam/decay
VariableV2*
shared_name *
dtype0*
_output_shapes
: *
	container *
shape: 
к
Adam/decay/AssignAssign
Adam/decayAdam/decay/initial_value*
validate_shape(*
_output_shapes
: *
use_locking(*
T0*
_class
loc:@Adam/decay
g
Adam/decay/readIdentity
Adam/decay*
T0*
_class
loc:@Adam/decay*
_output_shapes
: 
в
output_layer_targetPlaceholder*
dtype0*=
_output_shapes+
):'                           *2
shape):'                           
v
output_layer_sample_weightsPlaceholder*
shape:         *
dtype0*#
_output_shapes
:         
З
loss/output_layer_loss/subSuboutput_layer/addoutput_layer_target*
T0*4
_output_shapes"
 :                  (
|
loss/output_layer_loss/AbsAbsloss/output_layer_loss/sub*
T0*4
_output_shapes"
 :                  (
x
-loss/output_layer_loss/Mean/reduction_indicesConst*
valueB :
         *
dtype0*
_output_shapes
: 
╞
loss/output_layer_loss/MeanMeanloss/output_layer_loss/Abs-loss/output_layer_loss/Mean/reduction_indices*
T0*0
_output_shapes
:                  *
	keep_dims( *

Tidx0
y
/loss/output_layer_loss/Mean_1/reduction_indicesConst*
dtype0*
_output_shapes
:*
valueB:
╛
loss/output_layer_loss/Mean_1Meanloss/output_layer_loss/Mean/loss/output_layer_loss/Mean_1/reduction_indices*
T0*#
_output_shapes
:         *
	keep_dims( *

Tidx0
Л
loss/output_layer_loss/mulMulloss/output_layer_loss/Mean_1output_layer_sample_weights*
T0*#
_output_shapes
:         
f
!loss/output_layer_loss/NotEqual/yConst*
_output_shapes
: *
valueB
 *    *
dtype0
Щ
loss/output_layer_loss/NotEqualNotEqualoutput_layer_sample_weights!loss/output_layer_loss/NotEqual/y*#
_output_shapes
:         *
T0
С
loss/output_layer_loss/CastCastloss/output_layer_loss/NotEqual*
Truncate( *#
_output_shapes
:         *

DstT0*

SrcT0

f
loss/output_layer_loss/ConstConst*
dtype0*
_output_shapes
:*
valueB: 
Ю
loss/output_layer_loss/Mean_2Meanloss/output_layer_loss/Castloss/output_layer_loss/Const*
T0*
_output_shapes
: *
	keep_dims( *

Tidx0
Т
loss/output_layer_loss/truedivRealDivloss/output_layer_loss/mulloss/output_layer_loss/Mean_2*
T0*#
_output_shapes
:         
h
loss/output_layer_loss/Const_1Const*
valueB: *
dtype0*
_output_shapes
:
г
loss/output_layer_loss/Mean_3Meanloss/output_layer_loss/truedivloss/output_layer_loss/Const_1*
	keep_dims( *

Tidx0*
T0*
_output_shapes
: 
O

loss/mul/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
[
loss/mulMul
loss/mul/xloss/output_layer_loss/Mean_3*
T0*
_output_shapes
: 
М
metrics/mean_absolute_error/subSuboutput_layer/addoutput_layer_target*
T0*4
_output_shapes"
 :                  (
Ж
metrics/mean_absolute_error/AbsAbsmetrics/mean_absolute_error/sub*
T0*4
_output_shapes"
 :                  (
}
2metrics/mean_absolute_error/Mean/reduction_indicesConst*
_output_shapes
: *
valueB :
         *
dtype0
╒
 metrics/mean_absolute_error/MeanMeanmetrics/mean_absolute_error/Abs2metrics/mean_absolute_error/Mean/reduction_indices*
	keep_dims( *

Tidx0*
T0*0
_output_shapes
:                  
r
!metrics/mean_absolute_error/ConstConst*
valueB"       *
dtype0*
_output_shapes
:
н
"metrics/mean_absolute_error/Mean_1Mean metrics/mean_absolute_error/Mean!metrics/mean_absolute_error/Const*
T0*
_output_shapes
: *
	keep_dims( *

Tidx0
}
training/Adam/gradients/ShapeConst*
dtype0*
_output_shapes
: *
valueB *
_class
loc:@loss/mul
Г
!training/Adam/gradients/grad_ys_0Const*
valueB
 *  А?*
_class
loc:@loss/mul*
dtype0*
_output_shapes
: 
╢
training/Adam/gradients/FillFilltraining/Adam/gradients/Shape!training/Adam/gradients/grad_ys_0*
T0*

index_type0*
_class
loc:@loss/mul*
_output_shapes
: 
л
)training/Adam/gradients/loss/mul_grad/MulMultraining/Adam/gradients/Fillloss/output_layer_loss/Mean_3*
_output_shapes
: *
T0*
_class
loc:@loss/mul
Ъ
+training/Adam/gradients/loss/mul_grad/Mul_1Multraining/Adam/gradients/Fill
loss/mul/x*
T0*
_class
loc:@loss/mul*
_output_shapes
: 
─
Htraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Reshape/shapeConst*
_output_shapes
:*
valueB:*0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
dtype0
й
Btraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/ReshapeReshape+training/Adam/gradients/loss/mul_grad/Mul_1Htraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Reshape/shape*
T0*
Tshape0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
_output_shapes
:
╨
@training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/ShapeShapeloss/output_layer_loss/truediv*
_output_shapes
:*
T0*
out_type0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3
┐
?training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/TileTileBtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Reshape@training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Shape*

Tmultiples0*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3*#
_output_shapes
:         
╥
Btraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Shape_1Shapeloss/output_layer_loss/truediv*
_output_shapes
:*
T0*
out_type0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3
╖
Btraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Shape_2Const*
valueB *0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
dtype0*
_output_shapes
: 
╝
@training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/ConstConst*
valueB: *0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
dtype0*
_output_shapes
:
╜
?training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/ProdProdBtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Shape_1@training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Const*
_output_shapes
: *

Tidx0*
	keep_dims( *
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3
╛
Btraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Const_1Const*
valueB: *0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
dtype0*
_output_shapes
:
┴
Atraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Prod_1ProdBtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Shape_2Btraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Const_1*0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
_output_shapes
: *

Tidx0*
	keep_dims( *
T0
╕
Dtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Maximum/yConst*
value	B :*0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
dtype0*
_output_shapes
: 
й
Btraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/MaximumMaximumAtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Prod_1Dtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Maximum/y*
_output_shapes
: *
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3
з
Ctraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/floordivFloorDiv?training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/ProdBtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Maximum*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
_output_shapes
: 
■
?training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/CastCastCtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/floordiv*

SrcT0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3*
Truncate( *
_output_shapes
: *

DstT0
п
Btraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/truedivRealDiv?training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Tile?training/Adam/gradients/loss/output_layer_loss/Mean_3_grad/Cast*#
_output_shapes
:         *
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_3
╬
Atraining/Adam/gradients/loss/output_layer_loss/truediv_grad/ShapeShapeloss/output_layer_loss/mul*
out_type0*1
_class'
%#loc:@loss/output_layer_loss/truediv*
_output_shapes
:*
T0
╣
Ctraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Shape_1Const*
valueB *1
_class'
%#loc:@loss/output_layer_loss/truediv*
dtype0*
_output_shapes
: 
т
Qtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/BroadcastGradientArgsBroadcastGradientArgsAtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/ShapeCtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Shape_1*
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv*2
_output_shapes 
:         :         
Т
Ctraining/Adam/gradients/loss/output_layer_loss/truediv_grad/RealDivRealDivBtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/truedivloss/output_layer_loss/Mean_2*#
_output_shapes
:         *
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv
╤
?training/Adam/gradients/loss/output_layer_loss/truediv_grad/SumSumCtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/RealDivQtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/BroadcastGradientArgs*
_output_shapes
:*

Tidx0*
	keep_dims( *
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv
┴
Ctraining/Adam/gradients/loss/output_layer_loss/truediv_grad/ReshapeReshape?training/Adam/gradients/loss/output_layer_loss/truediv_grad/SumAtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Shape*#
_output_shapes
:         *
T0*
Tshape0*1
_class'
%#loc:@loss/output_layer_loss/truediv
├
?training/Adam/gradients/loss/output_layer_loss/truediv_grad/NegNegloss/output_layer_loss/mul*#
_output_shapes
:         *
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv
С
Etraining/Adam/gradients/loss/output_layer_loss/truediv_grad/RealDiv_1RealDiv?training/Adam/gradients/loss/output_layer_loss/truediv_grad/Negloss/output_layer_loss/Mean_2*
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv*#
_output_shapes
:         
Ч
Etraining/Adam/gradients/loss/output_layer_loss/truediv_grad/RealDiv_2RealDivEtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/RealDiv_1loss/output_layer_loss/Mean_2*
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv*#
_output_shapes
:         
▓
?training/Adam/gradients/loss/output_layer_loss/truediv_grad/mulMulBtraining/Adam/gradients/loss/output_layer_loss/Mean_3_grad/truedivEtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/RealDiv_2*
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv*#
_output_shapes
:         
╤
Atraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Sum_1Sum?training/Adam/gradients/loss/output_layer_loss/truediv_grad/mulStraining/Adam/gradients/loss/output_layer_loss/truediv_grad/BroadcastGradientArgs:1*
T0*1
_class'
%#loc:@loss/output_layer_loss/truediv*
_output_shapes
:*

Tidx0*
	keep_dims( 
║
Etraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Reshape_1ReshapeAtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Sum_1Ctraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Shape_1*
T0*
Tshape0*1
_class'
%#loc:@loss/output_layer_loss/truediv*
_output_shapes
: 
╔
=training/Adam/gradients/loss/output_layer_loss/mul_grad/ShapeShapeloss/output_layer_loss/Mean_1*
out_type0*-
_class#
!loc:@loss/output_layer_loss/mul*
_output_shapes
:*
T0
╔
?training/Adam/gradients/loss/output_layer_loss/mul_grad/Shape_1Shapeoutput_layer_sample_weights*
_output_shapes
:*
T0*
out_type0*-
_class#
!loc:@loss/output_layer_loss/mul
╥
Mtraining/Adam/gradients/loss/output_layer_loss/mul_grad/BroadcastGradientArgsBroadcastGradientArgs=training/Adam/gradients/loss/output_layer_loss/mul_grad/Shape?training/Adam/gradients/loss/output_layer_loss/mul_grad/Shape_1*
T0*-
_class#
!loc:@loss/output_layer_loss/mul*2
_output_shapes 
:         :         
Б
;training/Adam/gradients/loss/output_layer_loss/mul_grad/MulMulCtraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Reshapeoutput_layer_sample_weights*
T0*-
_class#
!loc:@loss/output_layer_loss/mul*#
_output_shapes
:         
╜
;training/Adam/gradients/loss/output_layer_loss/mul_grad/SumSum;training/Adam/gradients/loss/output_layer_loss/mul_grad/MulMtraining/Adam/gradients/loss/output_layer_loss/mul_grad/BroadcastGradientArgs*
T0*-
_class#
!loc:@loss/output_layer_loss/mul*
_output_shapes
:*

Tidx0*
	keep_dims( 
▒
?training/Adam/gradients/loss/output_layer_loss/mul_grad/ReshapeReshape;training/Adam/gradients/loss/output_layer_loss/mul_grad/Sum=training/Adam/gradients/loss/output_layer_loss/mul_grad/Shape*
T0*
Tshape0*-
_class#
!loc:@loss/output_layer_loss/mul*#
_output_shapes
:         
Е
=training/Adam/gradients/loss/output_layer_loss/mul_grad/Mul_1Mulloss/output_layer_loss/Mean_1Ctraining/Adam/gradients/loss/output_layer_loss/truediv_grad/Reshape*-
_class#
!loc:@loss/output_layer_loss/mul*#
_output_shapes
:         *
T0
├
=training/Adam/gradients/loss/output_layer_loss/mul_grad/Sum_1Sum=training/Adam/gradients/loss/output_layer_loss/mul_grad/Mul_1Otraining/Adam/gradients/loss/output_layer_loss/mul_grad/BroadcastGradientArgs:1*
T0*-
_class#
!loc:@loss/output_layer_loss/mul*
_output_shapes
:*

Tidx0*
	keep_dims( 
╖
Atraining/Adam/gradients/loss/output_layer_loss/mul_grad/Reshape_1Reshape=training/Adam/gradients/loss/output_layer_loss/mul_grad/Sum_1?training/Adam/gradients/loss/output_layer_loss/mul_grad/Shape_1*
T0*
Tshape0*-
_class#
!loc:@loss/output_layer_loss/mul*#
_output_shapes
:         
═
@training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/ShapeShapeloss/output_layer_loss/Mean*
_output_shapes
:*
T0*
out_type0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
│
?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/SizeConst*
value	B :*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
dtype0*
_output_shapes
: 
О
>training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/addAdd/loss/output_layer_loss/Mean_1/reduction_indices?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Size*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
_output_shapes
:
в
>training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/modFloorMod>training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/add?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Size*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
_output_shapes
:
╛
Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Shape_1Const*
valueB:*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
dtype0*
_output_shapes
:
║
Ftraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/range/startConst*
dtype0*
_output_shapes
: *
value	B : *0
_class&
$"loc:@loss/output_layer_loss/Mean_1
║
Ftraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/range/deltaConst*
value	B :*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
dtype0*
_output_shapes
: 
Ї
@training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/rangeRangeFtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/range/start?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/SizeFtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/range/delta*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
_output_shapes
:*

Tidx0
╣
Etraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Fill/valueConst*
dtype0*
_output_shapes
: *
value	B :*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
╗
?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/FillFillBtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Shape_1Etraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Fill/value*
T0*

index_type0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
_output_shapes
:
╛
Htraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/DynamicStitchDynamicStitch@training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/range>training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/mod@training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Shape?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Fill*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
N*
_output_shapes
:
╕
Dtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Maximum/yConst*
value	B :*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
dtype0*
_output_shapes
: 
┤
Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/MaximumMaximumHtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/DynamicStitchDtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Maximum/y*
_output_shapes
:*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
м
Ctraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/floordivFloorDiv@training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/ShapeBtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Maximum*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
_output_shapes
:
╙
Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/ReshapeReshape?training/Adam/gradients/loss/output_layer_loss/mul_grad/ReshapeHtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/DynamicStitch*
Tshape0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*0
_output_shapes
:                  *
T0
╧
?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/TileTileBtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/ReshapeCtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/floordiv*

Tmultiples0*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*0
_output_shapes
:                  
╧
Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Shape_2Shapeloss/output_layer_loss/Mean*
_output_shapes
:*
T0*
out_type0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
╤
Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Shape_3Shapeloss/output_layer_loss/Mean_1*
T0*
out_type0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
_output_shapes
:
╝
@training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/ConstConst*
valueB: *0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
dtype0*
_output_shapes
:
╜
?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/ProdProdBtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Shape_2@training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Const*
_output_shapes
: *

Tidx0*
	keep_dims( *
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
╛
Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Const_1Const*
valueB: *0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
dtype0*
_output_shapes
:
┴
Atraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Prod_1ProdBtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Shape_3Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Const_1*
_output_shapes
: *

Tidx0*
	keep_dims( *
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
║
Ftraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Maximum_1/yConst*
value	B :*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
dtype0*
_output_shapes
: 
н
Dtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Maximum_1MaximumAtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Prod_1Ftraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Maximum_1/y*
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
_output_shapes
: 
л
Etraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/floordiv_1FloorDiv?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/ProdDtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Maximum_1*
_output_shapes
: *
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
А
?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/CastCastEtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/floordiv_1*
_output_shapes
: *

DstT0*

SrcT0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1*
Truncate( 
╝
Btraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/truedivRealDiv?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Tile?training/Adam/gradients/loss/output_layer_loss/Mean_1_grad/Cast*0
_output_shapes
:                  *
T0*0
_class&
$"loc:@loss/output_layer_loss/Mean_1
╚
>training/Adam/gradients/loss/output_layer_loss/Mean_grad/ShapeShapeloss/output_layer_loss/Abs*
T0*
out_type0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
:
п
=training/Adam/gradients/loss/output_layer_loss/Mean_grad/SizeConst*
value	B :*.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
: 
В
<training/Adam/gradients/loss/output_layer_loss/Mean_grad/addAdd-loss/output_layer_loss/Mean/reduction_indices=training/Adam/gradients/loss/output_layer_loss/Mean_grad/Size*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
: 
Ц
<training/Adam/gradients/loss/output_layer_loss/Mean_grad/modFloorMod<training/Adam/gradients/loss/output_layer_loss/Mean_grad/add=training/Adam/gradients/loss/output_layer_loss/Mean_grad/Size*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
: 
│
@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape_1Const*
valueB *.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
: 
╢
Dtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/range/startConst*
dtype0*
_output_shapes
: *
value	B : *.
_class$
" loc:@loss/output_layer_loss/Mean
╢
Dtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/range/deltaConst*
value	B :*.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
: 
ъ
>training/Adam/gradients/loss/output_layer_loss/Mean_grad/rangeRangeDtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/range/start=training/Adam/gradients/loss/output_layer_loss/Mean_grad/SizeDtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/range/delta*

Tidx0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
:
╡
Ctraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Fill/valueConst*
value	B :*.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
: 
п
=training/Adam/gradients/loss/output_layer_loss/Mean_grad/FillFill@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape_1Ctraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Fill/value*
T0*

index_type0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
: 
▓
Ftraining/Adam/gradients/loss/output_layer_loss/Mean_grad/DynamicStitchDynamicStitch>training/Adam/gradients/loss/output_layer_loss/Mean_grad/range<training/Adam/gradients/loss/output_layer_loss/Mean_grad/mod>training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape=training/Adam/gradients/loss/output_layer_loss/Mean_grad/Fill*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
N*
_output_shapes
:
┤
Btraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Maximum/yConst*
value	B :*.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
: 
м
@training/Adam/gradients/loss/output_layer_loss/Mean_grad/MaximumMaximumFtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/DynamicStitchBtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Maximum/y*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
:
д
Atraining/Adam/gradients/loss/output_layer_loss/Mean_grad/floordivFloorDiv>training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Maximum*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
:
▌
@training/Adam/gradients/loss/output_layer_loss/Mean_grad/ReshapeReshapeBtraining/Adam/gradients/loss/output_layer_loss/Mean_1_grad/truedivFtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/DynamicStitch*
T0*
Tshape0*.
_class$
" loc:@loss/output_layer_loss/Mean*=
_output_shapes+
):'                           
╘
=training/Adam/gradients/loss/output_layer_loss/Mean_grad/TileTile@training/Adam/gradients/loss/output_layer_loss/Mean_grad/ReshapeAtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/floordiv*=
_output_shapes+
):'                           *

Tmultiples0*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean
╩
@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape_2Shapeloss/output_layer_loss/Abs*
_output_shapes
:*
T0*
out_type0*.
_class$
" loc:@loss/output_layer_loss/Mean
╦
@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape_3Shapeloss/output_layer_loss/Mean*
out_type0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
:*
T0
╕
>training/Adam/gradients/loss/output_layer_loss/Mean_grad/ConstConst*
valueB: *.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
:
╡
=training/Adam/gradients/loss/output_layer_loss/Mean_grad/ProdProd@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape_2>training/Adam/gradients/loss/output_layer_loss/Mean_grad/Const*
_output_shapes
: *

Tidx0*
	keep_dims( *
T0*.
_class$
" loc:@loss/output_layer_loss/Mean
║
@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Const_1Const*
valueB: *.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
:
╣
?training/Adam/gradients/loss/output_layer_loss/Mean_grad/Prod_1Prod@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Shape_3@training/Adam/gradients/loss/output_layer_loss/Mean_grad/Const_1*

Tidx0*
	keep_dims( *
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
: 
╢
Dtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Maximum_1/yConst*
value	B :*.
_class$
" loc:@loss/output_layer_loss/Mean*
dtype0*
_output_shapes
: 
е
Btraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Maximum_1Maximum?training/Adam/gradients/loss/output_layer_loss/Mean_grad/Prod_1Dtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Maximum_1/y*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
: 
г
Ctraining/Adam/gradients/loss/output_layer_loss/Mean_grad/floordiv_1FloorDiv=training/Adam/gradients/loss/output_layer_loss/Mean_grad/ProdBtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/Maximum_1*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*
_output_shapes
: 
·
=training/Adam/gradients/loss/output_layer_loss/Mean_grad/CastCastCtraining/Adam/gradients/loss/output_layer_loss/Mean_grad/floordiv_1*
_output_shapes
: *

DstT0*

SrcT0*.
_class$
" loc:@loss/output_layer_loss/Mean*
Truncate( 
╕
@training/Adam/gradients/loss/output_layer_loss/Mean_grad/truedivRealDiv=training/Adam/gradients/loss/output_layer_loss/Mean_grad/Tile=training/Adam/gradients/loss/output_layer_loss/Mean_grad/Cast*
T0*.
_class$
" loc:@loss/output_layer_loss/Mean*4
_output_shapes"
 :                  (
╬
<training/Adam/gradients/loss/output_layer_loss/Abs_grad/SignSignloss/output_layer_loss/sub*-
_class#
!loc:@loss/output_layer_loss/Abs*4
_output_shapes"
 :                  (*
T0
░
;training/Adam/gradients/loss/output_layer_loss/Abs_grad/mulMul@training/Adam/gradients/loss/output_layer_loss/Mean_grad/truediv<training/Adam/gradients/loss/output_layer_loss/Abs_grad/Sign*
T0*-
_class#
!loc:@loss/output_layer_loss/Abs*4
_output_shapes"
 :                  (
╝
=training/Adam/gradients/loss/output_layer_loss/sub_grad/ShapeShapeoutput_layer/add*
T0*
out_type0*-
_class#
!loc:@loss/output_layer_loss/sub*
_output_shapes
:
┴
?training/Adam/gradients/loss/output_layer_loss/sub_grad/Shape_1Shapeoutput_layer_target*
T0*
out_type0*-
_class#
!loc:@loss/output_layer_loss/sub*
_output_shapes
:
╥
Mtraining/Adam/gradients/loss/output_layer_loss/sub_grad/BroadcastGradientArgsBroadcastGradientArgs=training/Adam/gradients/loss/output_layer_loss/sub_grad/Shape?training/Adam/gradients/loss/output_layer_loss/sub_grad/Shape_1*
T0*-
_class#
!loc:@loss/output_layer_loss/sub*2
_output_shapes 
:         :         
╜
;training/Adam/gradients/loss/output_layer_loss/sub_grad/SumSum;training/Adam/gradients/loss/output_layer_loss/Abs_grad/mulMtraining/Adam/gradients/loss/output_layer_loss/sub_grad/BroadcastGradientArgs*
T0*-
_class#
!loc:@loss/output_layer_loss/sub*
_output_shapes
:*

Tidx0*
	keep_dims( 
╣
?training/Adam/gradients/loss/output_layer_loss/sub_grad/ReshapeReshape;training/Adam/gradients/loss/output_layer_loss/sub_grad/Sum=training/Adam/gradients/loss/output_layer_loss/sub_grad/Shape*
T0*
Tshape0*-
_class#
!loc:@loss/output_layer_loss/sub*+
_output_shapes
:         (
┴
=training/Adam/gradients/loss/output_layer_loss/sub_grad/Sum_1Sum;training/Adam/gradients/loss/output_layer_loss/Abs_grad/mulOtraining/Adam/gradients/loss/output_layer_loss/sub_grad/BroadcastGradientArgs:1*
T0*-
_class#
!loc:@loss/output_layer_loss/sub*
_output_shapes
:*

Tidx0*
	keep_dims( 
╙
;training/Adam/gradients/loss/output_layer_loss/sub_grad/NegNeg=training/Adam/gradients/loss/output_layer_loss/sub_grad/Sum_1*
T0*-
_class#
!loc:@loss/output_layer_loss/sub*
_output_shapes
:
╧
Atraining/Adam/gradients/loss/output_layer_loss/sub_grad/Reshape_1Reshape;training/Adam/gradients/loss/output_layer_loss/sub_grad/Neg?training/Adam/gradients/loss/output_layer_loss/sub_grad/Shape_1*
T0*
Tshape0*-
_class#
!loc:@loss/output_layer_loss/sub*=
_output_shapes+
):'                           
о
3training/Adam/gradients/output_layer/add_grad/ShapeShapeoutput_layer/Reshape_2*
out_type0*#
_class
loc:@output_layer/add*
_output_shapes
:*
T0
п
5training/Adam/gradients/output_layer/add_grad/Shape_1Const*
dtype0*
_output_shapes
:*!
valueB"      (   *#
_class
loc:@output_layer/add
к
Ctraining/Adam/gradients/output_layer/add_grad/BroadcastGradientArgsBroadcastGradientArgs3training/Adam/gradients/output_layer/add_grad/Shape5training/Adam/gradients/output_layer/add_grad/Shape_1*
T0*#
_class
loc:@output_layer/add*2
_output_shapes 
:         :         
г
1training/Adam/gradients/output_layer/add_grad/SumSum?training/Adam/gradients/loss/output_layer_loss/sub_grad/ReshapeCtraining/Adam/gradients/output_layer/add_grad/BroadcastGradientArgs*
T0*#
_class
loc:@output_layer/add*
_output_shapes
:*

Tidx0*
	keep_dims( 
С
5training/Adam/gradients/output_layer/add_grad/ReshapeReshape1training/Adam/gradients/output_layer/add_grad/Sum3training/Adam/gradients/output_layer/add_grad/Shape*
T0*
Tshape0*#
_class
loc:@output_layer/add*+
_output_shapes
:         (
з
3training/Adam/gradients/output_layer/add_grad/Sum_1Sum?training/Adam/gradients/loss/output_layer_loss/sub_grad/ReshapeEtraining/Adam/gradients/output_layer/add_grad/BroadcastGradientArgs:1*
T0*#
_class
loc:@output_layer/add*
_output_shapes
:*

Tidx0*
	keep_dims( 
О
7training/Adam/gradients/output_layer/add_grad/Reshape_1Reshape3training/Adam/gradients/output_layer/add_grad/Sum_15training/Adam/gradients/output_layer/add_grad/Shape_1*
T0*
Tshape0*#
_class
loc:@output_layer/add*"
_output_shapes
:(
╖
9training/Adam/gradients/output_layer/Reshape_2_grad/ShapeShapeoutput_layer/MatMul*
T0*
out_type0*)
_class
loc:@output_layer/Reshape_2*
_output_shapes
:
г
;training/Adam/gradients/output_layer/Reshape_2_grad/ReshapeReshape5training/Adam/gradients/output_layer/add_grad/Reshape9training/Adam/gradients/output_layer/Reshape_2_grad/Shape*
T0*
Tshape0*)
_class
loc:@output_layer/Reshape_2*'
_output_shapes
:         (
о
9training/Adam/gradients/output_layer/Reshape_3_grad/ShapeConst*
valueB:(*)
_class
loc:@output_layer/Reshape_3*
dtype0*
_output_shapes
:
Ш
;training/Adam/gradients/output_layer/Reshape_3_grad/ReshapeReshape7training/Adam/gradients/output_layer/add_grad/Reshape_19training/Adam/gradients/output_layer/Reshape_3_grad/Shape*
T0*
Tshape0*)
_class
loc:@output_layer/Reshape_3*
_output_shapes
:(
Ч
7training/Adam/gradients/output_layer/MatMul_grad/MatMulMatMul;training/Adam/gradients/output_layer/Reshape_2_grad/Reshapeoutput_layer/Reshape_1*
T0*&
_class
loc:@output_layer/MatMul*(
_output_shapes
:         А*
transpose_a( *
transpose_b(
О
9training/Adam/gradients/output_layer/MatMul_grad/MatMul_1MatMuloutput_layer/Reshape;training/Adam/gradients/output_layer/Reshape_2_grad/Reshape*&
_class
loc:@output_layer/MatMul*
_output_shapes
:	А(*
transpose_a(*
transpose_b( *
T0
м
7training/Adam/gradients/output_layer/Reshape_grad/ShapeShapedense_3/Relu*
T0*
out_type0*'
_class
loc:@output_layer/Reshape*
_output_shapes
:
д
9training/Adam/gradients/output_layer/Reshape_grad/ReshapeReshape7training/Adam/gradients/output_layer/MatMul_grad/MatMul7training/Adam/gradients/output_layer/Reshape_grad/Shape*
T0*
Tshape0*'
_class
loc:@output_layer/Reshape*,
_output_shapes
:         А
╡
9training/Adam/gradients/output_layer/Reshape_1_grad/ShapeConst*
valueB"А   (   *)
_class
loc:@output_layer/Reshape_1*
dtype0*
_output_shapes
:
Я
;training/Adam/gradients/output_layer/Reshape_1_grad/ReshapeReshape9training/Adam/gradients/output_layer/MatMul_grad/MatMul_19training/Adam/gradients/output_layer/Reshape_1_grad/Shape*
T0*
Tshape0*)
_class
loc:@output_layer/Reshape_1*
_output_shapes
:	А(
▀
2training/Adam/gradients/dense_3/Relu_grad/ReluGradReluGrad9training/Adam/gradients/output_layer/Reshape_grad/Reshapedense_3/Relu*
T0*
_class
loc:@dense_3/Relu*,
_output_shapes
:         А
╟
Etraining/Adam/gradients/output_layer/transpose_grad/InvertPermutationInvertPermutationoutput_layer/transpose/perm*
T0*)
_class
loc:@output_layer/transpose*
_output_shapes
:
░
=training/Adam/gradients/output_layer/transpose_grad/transpose	Transpose;training/Adam/gradients/output_layer/Reshape_1_grad/ReshapeEtraining/Adam/gradients/output_layer/transpose_grad/InvertPermutation*
T0*)
_class
loc:@output_layer/transpose*
_output_shapes
:	А(*
Tperm0
Я
.training/Adam/gradients/dense_3/add_grad/ShapeShapedense_3/Reshape_2*
T0*
out_type0*
_class
loc:@dense_3/add*
_output_shapes
:
е
0training/Adam/gradients/dense_3/add_grad/Shape_1Const*
dtype0*
_output_shapes
:*!
valueB"      А   *
_class
loc:@dense_3/add
Ц
>training/Adam/gradients/dense_3/add_grad/BroadcastGradientArgsBroadcastGradientArgs.training/Adam/gradients/dense_3/add_grad/Shape0training/Adam/gradients/dense_3/add_grad/Shape_1*2
_output_shapes 
:         :         *
T0*
_class
loc:@dense_3/add
З
,training/Adam/gradients/dense_3/add_grad/SumSum2training/Adam/gradients/dense_3/Relu_grad/ReluGrad>training/Adam/gradients/dense_3/add_grad/BroadcastGradientArgs*
T0*
_class
loc:@dense_3/add*
_output_shapes
:*

Tidx0*
	keep_dims( 
■
0training/Adam/gradients/dense_3/add_grad/ReshapeReshape,training/Adam/gradients/dense_3/add_grad/Sum.training/Adam/gradients/dense_3/add_grad/Shape*
Tshape0*
_class
loc:@dense_3/add*,
_output_shapes
:         А*
T0
Л
.training/Adam/gradients/dense_3/add_grad/Sum_1Sum2training/Adam/gradients/dense_3/Relu_grad/ReluGrad@training/Adam/gradients/dense_3/add_grad/BroadcastGradientArgs:1*
T0*
_class
loc:@dense_3/add*
_output_shapes
:*

Tidx0*
	keep_dims( 
√
2training/Adam/gradients/dense_3/add_grad/Reshape_1Reshape.training/Adam/gradients/dense_3/add_grad/Sum_10training/Adam/gradients/dense_3/add_grad/Shape_1*
Tshape0*
_class
loc:@dense_3/add*#
_output_shapes
:А*
T0
и
4training/Adam/gradients/dense_3/Reshape_2_grad/ShapeShapedense_3/MatMul*
_output_shapes
:*
T0*
out_type0*$
_class
loc:@dense_3/Reshape_2
Р
6training/Adam/gradients/dense_3/Reshape_2_grad/ReshapeReshape0training/Adam/gradients/dense_3/add_grad/Reshape4training/Adam/gradients/dense_3/Reshape_2_grad/Shape*
T0*
Tshape0*$
_class
loc:@dense_3/Reshape_2*(
_output_shapes
:         А
е
4training/Adam/gradients/dense_3/Reshape_3_grad/ShapeConst*
valueB:А*$
_class
loc:@dense_3/Reshape_3*
dtype0*
_output_shapes
:
Е
6training/Adam/gradients/dense_3/Reshape_3_grad/ReshapeReshape2training/Adam/gradients/dense_3/add_grad/Reshape_14training/Adam/gradients/dense_3/Reshape_3_grad/Shape*
T0*
Tshape0*$
_class
loc:@dense_3/Reshape_3*
_output_shapes	
:А
Г
2training/Adam/gradients/dense_3/MatMul_grad/MatMulMatMul6training/Adam/gradients/dense_3/Reshape_2_grad/Reshapedense_3/Reshape_1*
T0*!
_class
loc:@dense_3/MatMul*(
_output_shapes
:         А*
transpose_a( *
transpose_b(
√
4training/Adam/gradients/dense_3/MatMul_grad/MatMul_1MatMuldense_3/Reshape6training/Adam/gradients/dense_3/Reshape_2_grad/Reshape*
T0*!
_class
loc:@dense_3/MatMul* 
_output_shapes
:
АА*
transpose_a(*
transpose_b( 
в
2training/Adam/gradients/dense_3/Reshape_grad/ShapeShapedense_2/Relu*
T0*
out_type0*"
_class
loc:@dense_3/Reshape*
_output_shapes
:
Р
4training/Adam/gradients/dense_3/Reshape_grad/ReshapeReshape2training/Adam/gradients/dense_3/MatMul_grad/MatMul2training/Adam/gradients/dense_3/Reshape_grad/Shape*,
_output_shapes
:         А*
T0*
Tshape0*"
_class
loc:@dense_3/Reshape
л
4training/Adam/gradients/dense_3/Reshape_1_grad/ShapeConst*
valueB"А   А   *$
_class
loc:@dense_3/Reshape_1*
dtype0*
_output_shapes
:
М
6training/Adam/gradients/dense_3/Reshape_1_grad/ReshapeReshape4training/Adam/gradients/dense_3/MatMul_grad/MatMul_14training/Adam/gradients/dense_3/Reshape_1_grad/Shape*
Tshape0*$
_class
loc:@dense_3/Reshape_1* 
_output_shapes
:
АА*
T0
┌
2training/Adam/gradients/dense_2/Relu_grad/ReluGradReluGrad4training/Adam/gradients/dense_3/Reshape_grad/Reshapedense_2/Relu*
T0*
_class
loc:@dense_2/Relu*,
_output_shapes
:         А
╕
@training/Adam/gradients/dense_3/transpose_grad/InvertPermutationInvertPermutationdense_3/transpose/perm*
T0*$
_class
loc:@dense_3/transpose*
_output_shapes
:
Э
8training/Adam/gradients/dense_3/transpose_grad/transpose	Transpose6training/Adam/gradients/dense_3/Reshape_1_grad/Reshape@training/Adam/gradients/dense_3/transpose_grad/InvertPermutation*
T0*$
_class
loc:@dense_3/transpose* 
_output_shapes
:
АА*
Tperm0
Я
.training/Adam/gradients/dense_2/add_grad/ShapeShapedense_2/Reshape_2*
T0*
out_type0*
_class
loc:@dense_2/add*
_output_shapes
:
е
0training/Adam/gradients/dense_2/add_grad/Shape_1Const*!
valueB"      А   *
_class
loc:@dense_2/add*
dtype0*
_output_shapes
:
Ц
>training/Adam/gradients/dense_2/add_grad/BroadcastGradientArgsBroadcastGradientArgs.training/Adam/gradients/dense_2/add_grad/Shape0training/Adam/gradients/dense_2/add_grad/Shape_1*
T0*
_class
loc:@dense_2/add*2
_output_shapes 
:         :         
З
,training/Adam/gradients/dense_2/add_grad/SumSum2training/Adam/gradients/dense_2/Relu_grad/ReluGrad>training/Adam/gradients/dense_2/add_grad/BroadcastGradientArgs*

Tidx0*
	keep_dims( *
T0*
_class
loc:@dense_2/add*
_output_shapes
:
■
0training/Adam/gradients/dense_2/add_grad/ReshapeReshape,training/Adam/gradients/dense_2/add_grad/Sum.training/Adam/gradients/dense_2/add_grad/Shape*
T0*
Tshape0*
_class
loc:@dense_2/add*,
_output_shapes
:         А
Л
.training/Adam/gradients/dense_2/add_grad/Sum_1Sum2training/Adam/gradients/dense_2/Relu_grad/ReluGrad@training/Adam/gradients/dense_2/add_grad/BroadcastGradientArgs:1*
T0*
_class
loc:@dense_2/add*
_output_shapes
:*

Tidx0*
	keep_dims( 
√
2training/Adam/gradients/dense_2/add_grad/Reshape_1Reshape.training/Adam/gradients/dense_2/add_grad/Sum_10training/Adam/gradients/dense_2/add_grad/Shape_1*
T0*
Tshape0*
_class
loc:@dense_2/add*#
_output_shapes
:А
и
4training/Adam/gradients/dense_2/Reshape_2_grad/ShapeShapedense_2/MatMul*
T0*
out_type0*$
_class
loc:@dense_2/Reshape_2*
_output_shapes
:
Р
6training/Adam/gradients/dense_2/Reshape_2_grad/ReshapeReshape0training/Adam/gradients/dense_2/add_grad/Reshape4training/Adam/gradients/dense_2/Reshape_2_grad/Shape*(
_output_shapes
:         А*
T0*
Tshape0*$
_class
loc:@dense_2/Reshape_2
е
4training/Adam/gradients/dense_2/Reshape_3_grad/ShapeConst*
valueB:А*$
_class
loc:@dense_2/Reshape_3*
dtype0*
_output_shapes
:
Е
6training/Adam/gradients/dense_2/Reshape_3_grad/ReshapeReshape2training/Adam/gradients/dense_2/add_grad/Reshape_14training/Adam/gradients/dense_2/Reshape_3_grad/Shape*
T0*
Tshape0*$
_class
loc:@dense_2/Reshape_3*
_output_shapes	
:А
Г
2training/Adam/gradients/dense_2/MatMul_grad/MatMulMatMul6training/Adam/gradients/dense_2/Reshape_2_grad/Reshapedense_2/Reshape_1*(
_output_shapes
:         А*
transpose_a( *
transpose_b(*
T0*!
_class
loc:@dense_2/MatMul
√
4training/Adam/gradients/dense_2/MatMul_grad/MatMul_1MatMuldense_2/Reshape6training/Adam/gradients/dense_2/Reshape_2_grad/Reshape*
transpose_b( *
T0*!
_class
loc:@dense_2/MatMul* 
_output_shapes
:
АА*
transpose_a(
в
2training/Adam/gradients/dense_2/Reshape_grad/ShapeShapedense_1/Relu*
T0*
out_type0*"
_class
loc:@dense_2/Reshape*
_output_shapes
:
Р
4training/Adam/gradients/dense_2/Reshape_grad/ReshapeReshape2training/Adam/gradients/dense_2/MatMul_grad/MatMul2training/Adam/gradients/dense_2/Reshape_grad/Shape*
T0*
Tshape0*"
_class
loc:@dense_2/Reshape*,
_output_shapes
:         А
л
4training/Adam/gradients/dense_2/Reshape_1_grad/ShapeConst*
dtype0*
_output_shapes
:*
valueB"А   А   *$
_class
loc:@dense_2/Reshape_1
М
6training/Adam/gradients/dense_2/Reshape_1_grad/ReshapeReshape4training/Adam/gradients/dense_2/MatMul_grad/MatMul_14training/Adam/gradients/dense_2/Reshape_1_grad/Shape* 
_output_shapes
:
АА*
T0*
Tshape0*$
_class
loc:@dense_2/Reshape_1
┌
2training/Adam/gradients/dense_1/Relu_grad/ReluGradReluGrad4training/Adam/gradients/dense_2/Reshape_grad/Reshapedense_1/Relu*
T0*
_class
loc:@dense_1/Relu*,
_output_shapes
:         А
╕
@training/Adam/gradients/dense_2/transpose_grad/InvertPermutationInvertPermutationdense_2/transpose/perm*
T0*$
_class
loc:@dense_2/transpose*
_output_shapes
:
Э
8training/Adam/gradients/dense_2/transpose_grad/transpose	Transpose6training/Adam/gradients/dense_2/Reshape_1_grad/Reshape@training/Adam/gradients/dense_2/transpose_grad/InvertPermutation* 
_output_shapes
:
АА*
Tperm0*
T0*$
_class
loc:@dense_2/transpose
Я
.training/Adam/gradients/dense_1/add_grad/ShapeShapedense_1/Reshape_2*
T0*
out_type0*
_class
loc:@dense_1/add*
_output_shapes
:
е
0training/Adam/gradients/dense_1/add_grad/Shape_1Const*
dtype0*
_output_shapes
:*!
valueB"      А   *
_class
loc:@dense_1/add
Ц
>training/Adam/gradients/dense_1/add_grad/BroadcastGradientArgsBroadcastGradientArgs.training/Adam/gradients/dense_1/add_grad/Shape0training/Adam/gradients/dense_1/add_grad/Shape_1*
T0*
_class
loc:@dense_1/add*2
_output_shapes 
:         :         
З
,training/Adam/gradients/dense_1/add_grad/SumSum2training/Adam/gradients/dense_1/Relu_grad/ReluGrad>training/Adam/gradients/dense_1/add_grad/BroadcastGradientArgs*
_output_shapes
:*

Tidx0*
	keep_dims( *
T0*
_class
loc:@dense_1/add
■
0training/Adam/gradients/dense_1/add_grad/ReshapeReshape,training/Adam/gradients/dense_1/add_grad/Sum.training/Adam/gradients/dense_1/add_grad/Shape*
T0*
Tshape0*
_class
loc:@dense_1/add*,
_output_shapes
:         А
Л
.training/Adam/gradients/dense_1/add_grad/Sum_1Sum2training/Adam/gradients/dense_1/Relu_grad/ReluGrad@training/Adam/gradients/dense_1/add_grad/BroadcastGradientArgs:1*
_output_shapes
:*

Tidx0*
	keep_dims( *
T0*
_class
loc:@dense_1/add
√
2training/Adam/gradients/dense_1/add_grad/Reshape_1Reshape.training/Adam/gradients/dense_1/add_grad/Sum_10training/Adam/gradients/dense_1/add_grad/Shape_1*#
_output_shapes
:А*
T0*
Tshape0*
_class
loc:@dense_1/add
и
4training/Adam/gradients/dense_1/Reshape_2_grad/ShapeShapedense_1/MatMul*
T0*
out_type0*$
_class
loc:@dense_1/Reshape_2*
_output_shapes
:
Р
6training/Adam/gradients/dense_1/Reshape_2_grad/ReshapeReshape0training/Adam/gradients/dense_1/add_grad/Reshape4training/Adam/gradients/dense_1/Reshape_2_grad/Shape*
T0*
Tshape0*$
_class
loc:@dense_1/Reshape_2*(
_output_shapes
:         А
е
4training/Adam/gradients/dense_1/Reshape_3_grad/ShapeConst*
valueB:А*$
_class
loc:@dense_1/Reshape_3*
dtype0*
_output_shapes
:
Е
6training/Adam/gradients/dense_1/Reshape_3_grad/ReshapeReshape2training/Adam/gradients/dense_1/add_grad/Reshape_14training/Adam/gradients/dense_1/Reshape_3_grad/Shape*
_output_shapes	
:А*
T0*
Tshape0*$
_class
loc:@dense_1/Reshape_3
Г
2training/Adam/gradients/dense_1/MatMul_grad/MatMulMatMul6training/Adam/gradients/dense_1/Reshape_2_grad/Reshapedense_1/Reshape_1*
T0*!
_class
loc:@dense_1/MatMul*(
_output_shapes
:         А*
transpose_a( *
transpose_b(
√
4training/Adam/gradients/dense_1/MatMul_grad/MatMul_1MatMuldense_1/Reshape6training/Adam/gradients/dense_1/Reshape_2_grad/Reshape*
T0*!
_class
loc:@dense_1/MatMul* 
_output_shapes
:
АА*
transpose_a(*
transpose_b( 
ж
2training/Adam/gradients/dense_1/Reshape_grad/ShapeShapeinput_layer/Relu*
T0*
out_type0*"
_class
loc:@dense_1/Reshape*
_output_shapes
:
Р
4training/Adam/gradients/dense_1/Reshape_grad/ReshapeReshape2training/Adam/gradients/dense_1/MatMul_grad/MatMul2training/Adam/gradients/dense_1/Reshape_grad/Shape*
T0*
Tshape0*"
_class
loc:@dense_1/Reshape*,
_output_shapes
:         А
л
4training/Adam/gradients/dense_1/Reshape_1_grad/ShapeConst*
_output_shapes
:*
valueB"   А   *$
_class
loc:@dense_1/Reshape_1*
dtype0
М
6training/Adam/gradients/dense_1/Reshape_1_grad/ReshapeReshape4training/Adam/gradients/dense_1/MatMul_grad/MatMul_14training/Adam/gradients/dense_1/Reshape_1_grad/Shape*
T0*
Tshape0*$
_class
loc:@dense_1/Reshape_1* 
_output_shapes
:
АА
ц
6training/Adam/gradients/input_layer/Relu_grad/ReluGradReluGrad4training/Adam/gradients/dense_1/Reshape_grad/Reshapeinput_layer/Relu*
T0*#
_class
loc:@input_layer/Relu*,
_output_shapes
:         А
╕
@training/Adam/gradients/dense_1/transpose_grad/InvertPermutationInvertPermutationdense_1/transpose/perm*
T0*$
_class
loc:@dense_1/transpose*
_output_shapes
:
Э
8training/Adam/gradients/dense_1/transpose_grad/transpose	Transpose6training/Adam/gradients/dense_1/Reshape_1_grad/Reshape@training/Adam/gradients/dense_1/transpose_grad/InvertPermutation*
Tperm0*
T0*$
_class
loc:@dense_1/transpose* 
_output_shapes
:
АА
л
2training/Adam/gradients/input_layer/add_grad/ShapeShapeinput_layer/Reshape_2*
T0*
out_type0*"
_class
loc:@input_layer/add*
_output_shapes
:
н
4training/Adam/gradients/input_layer/add_grad/Shape_1Const*!
valueB"         *"
_class
loc:@input_layer/add*
dtype0*
_output_shapes
:
ж
Btraining/Adam/gradients/input_layer/add_grad/BroadcastGradientArgsBroadcastGradientArgs2training/Adam/gradients/input_layer/add_grad/Shape4training/Adam/gradients/input_layer/add_grad/Shape_1*
T0*"
_class
loc:@input_layer/add*2
_output_shapes 
:         :         
Ч
0training/Adam/gradients/input_layer/add_grad/SumSum6training/Adam/gradients/input_layer/Relu_grad/ReluGradBtraining/Adam/gradients/input_layer/add_grad/BroadcastGradientArgs*
T0*"
_class
loc:@input_layer/add*
_output_shapes
:*

Tidx0*
	keep_dims( 
О
4training/Adam/gradients/input_layer/add_grad/ReshapeReshape0training/Adam/gradients/input_layer/add_grad/Sum2training/Adam/gradients/input_layer/add_grad/Shape*
T0*
Tshape0*"
_class
loc:@input_layer/add*,
_output_shapes
:         А
Ы
2training/Adam/gradients/input_layer/add_grad/Sum_1Sum6training/Adam/gradients/input_layer/Relu_grad/ReluGradDtraining/Adam/gradients/input_layer/add_grad/BroadcastGradientArgs:1*

Tidx0*
	keep_dims( *
T0*"
_class
loc:@input_layer/add*
_output_shapes
:
Л
6training/Adam/gradients/input_layer/add_grad/Reshape_1Reshape2training/Adam/gradients/input_layer/add_grad/Sum_14training/Adam/gradients/input_layer/add_grad/Shape_1*
Tshape0*"
_class
loc:@input_layer/add*#
_output_shapes
:А*
T0
┤
8training/Adam/gradients/input_layer/Reshape_2_grad/ShapeShapeinput_layer/MatMul*
out_type0*(
_class
loc:@input_layer/Reshape_2*
_output_shapes
:*
T0
а
:training/Adam/gradients/input_layer/Reshape_2_grad/ReshapeReshape4training/Adam/gradients/input_layer/add_grad/Reshape8training/Adam/gradients/input_layer/Reshape_2_grad/Shape*
T0*
Tshape0*(
_class
loc:@input_layer/Reshape_2*(
_output_shapes
:         А
н
8training/Adam/gradients/input_layer/Reshape_3_grad/ShapeConst*
valueB:А*(
_class
loc:@input_layer/Reshape_3*
dtype0*
_output_shapes
:
Х
:training/Adam/gradients/input_layer/Reshape_3_grad/ReshapeReshape6training/Adam/gradients/input_layer/add_grad/Reshape_18training/Adam/gradients/input_layer/Reshape_3_grad/Shape*
T0*
Tshape0*(
_class
loc:@input_layer/Reshape_3*
_output_shapes	
:А
Т
6training/Adam/gradients/input_layer/MatMul_grad/MatMulMatMul:training/Adam/gradients/input_layer/Reshape_2_grad/Reshapeinput_layer/Reshape_1*
transpose_b(*
T0*%
_class
loc:@input_layer/MatMul*'
_output_shapes
:         z*
transpose_a( 
К
8training/Adam/gradients/input_layer/MatMul_grad/MatMul_1MatMulinput_layer/Reshape:training/Adam/gradients/input_layer/Reshape_2_grad/Reshape*
T0*%
_class
loc:@input_layer/MatMul*
_output_shapes
:	zА*
transpose_a(*
transpose_b( 
│
8training/Adam/gradients/input_layer/Reshape_1_grad/ShapeConst*
valueB"z      *(
_class
loc:@input_layer/Reshape_1*
dtype0*
_output_shapes
:
Ы
:training/Adam/gradients/input_layer/Reshape_1_grad/ReshapeReshape8training/Adam/gradients/input_layer/MatMul_grad/MatMul_18training/Adam/gradients/input_layer/Reshape_1_grad/Shape*
_output_shapes
:	zА*
T0*
Tshape0*(
_class
loc:@input_layer/Reshape_1
─
Dtraining/Adam/gradients/input_layer/transpose_grad/InvertPermutationInvertPermutationinput_layer/transpose/perm*(
_class
loc:@input_layer/transpose*
_output_shapes
:*
T0
м
<training/Adam/gradients/input_layer/transpose_grad/transpose	Transpose:training/Adam/gradients/input_layer/Reshape_1_grad/ReshapeDtraining/Adam/gradients/input_layer/transpose_grad/InvertPermutation*
T0*(
_class
loc:@input_layer/transpose*
_output_shapes
:	zА*
Tperm0
_
training/Adam/AssignAdd/valueConst*
dtype0	*
_output_shapes
: *
value	B	 R
м
training/Adam/AssignAdd	AssignAddAdam/iterationstraining/Adam/AssignAdd/value*"
_class
loc:@Adam/iterations*
_output_shapes
: *
use_locking( *
T0	
p
training/Adam/CastCastAdam/iterations/read*
Truncate( *
_output_shapes
: *

DstT0*

SrcT0	
X
training/Adam/add/yConst*
dtype0*
_output_shapes
: *
valueB
 *  А?
b
training/Adam/addAddtraining/Adam/Casttraining/Adam/add/y*
T0*
_output_shapes
: 
^
training/Adam/PowPowAdam/beta_2/readtraining/Adam/add*
_output_shapes
: *
T0
X
training/Adam/sub/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
a
training/Adam/subSubtraining/Adam/sub/xtraining/Adam/Pow*
T0*
_output_shapes
: 
X
training/Adam/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Z
training/Adam/Const_1Const*
_output_shapes
: *
valueB
 *  А*
dtype0
y
#training/Adam/clip_by_value/MinimumMinimumtraining/Adam/subtraining/Adam/Const_1*
T0*
_output_shapes
: 
Б
training/Adam/clip_by_valueMaximum#training/Adam/clip_by_value/Minimumtraining/Adam/Const*
_output_shapes
: *
T0
X
training/Adam/SqrtSqrttraining/Adam/clip_by_value*
_output_shapes
: *
T0
`
training/Adam/Pow_1PowAdam/beta_1/readtraining/Adam/add*
T0*
_output_shapes
: 
Z
training/Adam/sub_1/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
g
training/Adam/sub_1Subtraining/Adam/sub_1/xtraining/Adam/Pow_1*
_output_shapes
: *
T0
j
training/Adam/truedivRealDivtraining/Adam/Sqrttraining/Adam/sub_1*
_output_shapes
: *
T0
^
training/Adam/mulMulAdam/lr/readtraining/Adam/truediv*
T0*
_output_shapes
: 
t
#training/Adam/zeros/shape_as_tensorConst*
valueB"z      *
dtype0*
_output_shapes
:
^
training/Adam/zeros/ConstConst*
_output_shapes
: *
valueB
 *    *
dtype0
Ч
training/Adam/zerosFill#training/Adam/zeros/shape_as_tensortraining/Adam/zeros/Const*
T0*

index_type0*
_output_shapes
:	zА
М
training/Adam/Variable
VariableV2*
shared_name *
dtype0*
_output_shapes
:	zА*
	container *
shape:	zА
╥
training/Adam/Variable/AssignAssigntraining/Adam/Variabletraining/Adam/zeros*)
_class
loc:@training/Adam/Variable*
validate_shape(*
_output_shapes
:	zА*
use_locking(*
T0
Ф
training/Adam/Variable/readIdentitytraining/Adam/Variable*
T0*)
_class
loc:@training/Adam/Variable*
_output_shapes
:	zА
d
training/Adam/zeros_1Const*
valueBА*    *
dtype0*
_output_shapes	
:А
Ж
training/Adam/Variable_1
VariableV2*
shared_name *
dtype0*
_output_shapes	
:А*
	container *
shape:А
╓
training/Adam/Variable_1/AssignAssigntraining/Adam/Variable_1training/Adam/zeros_1*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_1*
validate_shape(*
_output_shapes	
:А
Ц
training/Adam/Variable_1/readIdentitytraining/Adam/Variable_1*
T0*+
_class!
loc:@training/Adam/Variable_1*
_output_shapes	
:А
v
%training/Adam/zeros_2/shape_as_tensorConst*
valueB"   А   *
dtype0*
_output_shapes
:
`
training/Adam/zeros_2/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ю
training/Adam/zeros_2Fill%training/Adam/zeros_2/shape_as_tensortraining/Adam/zeros_2/Const*

index_type0* 
_output_shapes
:
АА*
T0
Р
training/Adam/Variable_2
VariableV2*
dtype0* 
_output_shapes
:
АА*
	container *
shape:
АА*
shared_name 
█
training/Adam/Variable_2/AssignAssigntraining/Adam/Variable_2training/Adam/zeros_2*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_2
Ы
training/Adam/Variable_2/readIdentitytraining/Adam/Variable_2* 
_output_shapes
:
АА*
T0*+
_class!
loc:@training/Adam/Variable_2
d
training/Adam/zeros_3Const*
valueBА*    *
dtype0*
_output_shapes	
:А
Ж
training/Adam/Variable_3
VariableV2*
shape:А*
shared_name *
dtype0*
_output_shapes	
:А*
	container 
╓
training/Adam/Variable_3/AssignAssigntraining/Adam/Variable_3training/Adam/zeros_3*+
_class!
loc:@training/Adam/Variable_3*
validate_shape(*
_output_shapes	
:А*
use_locking(*
T0
Ц
training/Adam/Variable_3/readIdentitytraining/Adam/Variable_3*
_output_shapes	
:А*
T0*+
_class!
loc:@training/Adam/Variable_3
v
%training/Adam/zeros_4/shape_as_tensorConst*
valueB"А   А   *
dtype0*
_output_shapes
:
`
training/Adam/zeros_4/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ю
training/Adam/zeros_4Fill%training/Adam/zeros_4/shape_as_tensortraining/Adam/zeros_4/Const*
T0*

index_type0* 
_output_shapes
:
АА
Р
training/Adam/Variable_4
VariableV2*
shape:
АА*
shared_name *
dtype0* 
_output_shapes
:
АА*
	container 
█
training/Adam/Variable_4/AssignAssigntraining/Adam/Variable_4training/Adam/zeros_4*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_4*
validate_shape(* 
_output_shapes
:
АА
Ы
training/Adam/Variable_4/readIdentitytraining/Adam/Variable_4* 
_output_shapes
:
АА*
T0*+
_class!
loc:@training/Adam/Variable_4
d
training/Adam/zeros_5Const*
_output_shapes	
:А*
valueBА*    *
dtype0
Ж
training/Adam/Variable_5
VariableV2*
dtype0*
_output_shapes	
:А*
	container *
shape:А*
shared_name 
╓
training/Adam/Variable_5/AssignAssigntraining/Adam/Variable_5training/Adam/zeros_5*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_5*
validate_shape(*
_output_shapes	
:А
Ц
training/Adam/Variable_5/readIdentitytraining/Adam/Variable_5*
T0*+
_class!
loc:@training/Adam/Variable_5*
_output_shapes	
:А
v
%training/Adam/zeros_6/shape_as_tensorConst*
_output_shapes
:*
valueB"А   А   *
dtype0
`
training/Adam/zeros_6/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ю
training/Adam/zeros_6Fill%training/Adam/zeros_6/shape_as_tensortraining/Adam/zeros_6/Const*

index_type0* 
_output_shapes
:
АА*
T0
Р
training/Adam/Variable_6
VariableV2*
dtype0* 
_output_shapes
:
АА*
	container *
shape:
АА*
shared_name 
█
training/Adam/Variable_6/AssignAssigntraining/Adam/Variable_6training/Adam/zeros_6*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_6*
validate_shape(* 
_output_shapes
:
АА
Ы
training/Adam/Variable_6/readIdentitytraining/Adam/Variable_6*
T0*+
_class!
loc:@training/Adam/Variable_6* 
_output_shapes
:
АА
d
training/Adam/zeros_7Const*
valueBА*    *
dtype0*
_output_shapes	
:А
Ж
training/Adam/Variable_7
VariableV2*
shared_name *
dtype0*
_output_shapes	
:А*
	container *
shape:А
╓
training/Adam/Variable_7/AssignAssigntraining/Adam/Variable_7training/Adam/zeros_7*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_7*
validate_shape(*
_output_shapes	
:А
Ц
training/Adam/Variable_7/readIdentitytraining/Adam/Variable_7*
_output_shapes	
:А*
T0*+
_class!
loc:@training/Adam/Variable_7
v
%training/Adam/zeros_8/shape_as_tensorConst*
_output_shapes
:*
valueB"А   (   *
dtype0
`
training/Adam/zeros_8/ConstConst*
dtype0*
_output_shapes
: *
valueB
 *    
Э
training/Adam/zeros_8Fill%training/Adam/zeros_8/shape_as_tensortraining/Adam/zeros_8/Const*
T0*

index_type0*
_output_shapes
:	А(
О
training/Adam/Variable_8
VariableV2*
_output_shapes
:	А(*
	container *
shape:	А(*
shared_name *
dtype0
┌
training/Adam/Variable_8/AssignAssigntraining/Adam/Variable_8training/Adam/zeros_8*
_output_shapes
:	А(*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_8*
validate_shape(
Ъ
training/Adam/Variable_8/readIdentitytraining/Adam/Variable_8*
T0*+
_class!
loc:@training/Adam/Variable_8*
_output_shapes
:	А(
b
training/Adam/zeros_9Const*
_output_shapes
:(*
valueB(*    *
dtype0
Д
training/Adam/Variable_9
VariableV2*
dtype0*
_output_shapes
:(*
	container *
shape:(*
shared_name 
╒
training/Adam/Variable_9/AssignAssigntraining/Adam/Variable_9training/Adam/zeros_9*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_9*
validate_shape(*
_output_shapes
:(
Х
training/Adam/Variable_9/readIdentitytraining/Adam/Variable_9*
T0*+
_class!
loc:@training/Adam/Variable_9*
_output_shapes
:(
w
&training/Adam/zeros_10/shape_as_tensorConst*
valueB"z      *
dtype0*
_output_shapes
:
a
training/Adam/zeros_10/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
а
training/Adam/zeros_10Fill&training/Adam/zeros_10/shape_as_tensortraining/Adam/zeros_10/Const*
T0*

index_type0*
_output_shapes
:	zА
П
training/Adam/Variable_10
VariableV2*
dtype0*
_output_shapes
:	zА*
	container *
shape:	zА*
shared_name 
▐
 training/Adam/Variable_10/AssignAssigntraining/Adam/Variable_10training/Adam/zeros_10*
T0*,
_class"
 loc:@training/Adam/Variable_10*
validate_shape(*
_output_shapes
:	zА*
use_locking(
Э
training/Adam/Variable_10/readIdentitytraining/Adam/Variable_10*
T0*,
_class"
 loc:@training/Adam/Variable_10*
_output_shapes
:	zА
e
training/Adam/zeros_11Const*
valueBА*    *
dtype0*
_output_shapes	
:А
З
training/Adam/Variable_11
VariableV2*
dtype0*
_output_shapes	
:А*
	container *
shape:А*
shared_name 
┌
 training/Adam/Variable_11/AssignAssigntraining/Adam/Variable_11training/Adam/zeros_11*
_output_shapes	
:А*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_11*
validate_shape(
Щ
training/Adam/Variable_11/readIdentitytraining/Adam/Variable_11*
_output_shapes	
:А*
T0*,
_class"
 loc:@training/Adam/Variable_11
w
&training/Adam/zeros_12/shape_as_tensorConst*
valueB"   А   *
dtype0*
_output_shapes
:
a
training/Adam/zeros_12/ConstConst*
dtype0*
_output_shapes
: *
valueB
 *    
б
training/Adam/zeros_12Fill&training/Adam/zeros_12/shape_as_tensortraining/Adam/zeros_12/Const* 
_output_shapes
:
АА*
T0*

index_type0
С
training/Adam/Variable_12
VariableV2*
shape:
АА*
shared_name *
dtype0* 
_output_shapes
:
АА*
	container 
▀
 training/Adam/Variable_12/AssignAssigntraining/Adam/Variable_12training/Adam/zeros_12*
T0*,
_class"
 loc:@training/Adam/Variable_12*
validate_shape(* 
_output_shapes
:
АА*
use_locking(
Ю
training/Adam/Variable_12/readIdentitytraining/Adam/Variable_12* 
_output_shapes
:
АА*
T0*,
_class"
 loc:@training/Adam/Variable_12
e
training/Adam/zeros_13Const*
valueBА*    *
dtype0*
_output_shapes	
:А
З
training/Adam/Variable_13
VariableV2*
dtype0*
_output_shapes	
:А*
	container *
shape:А*
shared_name 
┌
 training/Adam/Variable_13/AssignAssigntraining/Adam/Variable_13training/Adam/zeros_13*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_13*
validate_shape(*
_output_shapes	
:А
Щ
training/Adam/Variable_13/readIdentitytraining/Adam/Variable_13*
_output_shapes	
:А*
T0*,
_class"
 loc:@training/Adam/Variable_13
w
&training/Adam/zeros_14/shape_as_tensorConst*
_output_shapes
:*
valueB"А   А   *
dtype0
a
training/Adam/zeros_14/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
б
training/Adam/zeros_14Fill&training/Adam/zeros_14/shape_as_tensortraining/Adam/zeros_14/Const*
T0*

index_type0* 
_output_shapes
:
АА
С
training/Adam/Variable_14
VariableV2*
shape:
АА*
shared_name *
dtype0* 
_output_shapes
:
АА*
	container 
▀
 training/Adam/Variable_14/AssignAssigntraining/Adam/Variable_14training/Adam/zeros_14*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_14*
validate_shape(* 
_output_shapes
:
АА
Ю
training/Adam/Variable_14/readIdentitytraining/Adam/Variable_14* 
_output_shapes
:
АА*
T0*,
_class"
 loc:@training/Adam/Variable_14
e
training/Adam/zeros_15Const*
valueBА*    *
dtype0*
_output_shapes	
:А
З
training/Adam/Variable_15
VariableV2*
shape:А*
shared_name *
dtype0*
_output_shapes	
:А*
	container 
┌
 training/Adam/Variable_15/AssignAssigntraining/Adam/Variable_15training/Adam/zeros_15*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_15*
validate_shape(*
_output_shapes	
:А
Щ
training/Adam/Variable_15/readIdentitytraining/Adam/Variable_15*
_output_shapes	
:А*
T0*,
_class"
 loc:@training/Adam/Variable_15
w
&training/Adam/zeros_16/shape_as_tensorConst*
dtype0*
_output_shapes
:*
valueB"А   А   
a
training/Adam/zeros_16/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
б
training/Adam/zeros_16Fill&training/Adam/zeros_16/shape_as_tensortraining/Adam/zeros_16/Const*

index_type0* 
_output_shapes
:
АА*
T0
С
training/Adam/Variable_16
VariableV2* 
_output_shapes
:
АА*
	container *
shape:
АА*
shared_name *
dtype0
▀
 training/Adam/Variable_16/AssignAssigntraining/Adam/Variable_16training/Adam/zeros_16*,
_class"
 loc:@training/Adam/Variable_16*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0
Ю
training/Adam/Variable_16/readIdentitytraining/Adam/Variable_16*
T0*,
_class"
 loc:@training/Adam/Variable_16* 
_output_shapes
:
АА
e
training/Adam/zeros_17Const*
valueBА*    *
dtype0*
_output_shapes	
:А
З
training/Adam/Variable_17
VariableV2*
shared_name *
dtype0*
_output_shapes	
:А*
	container *
shape:А
┌
 training/Adam/Variable_17/AssignAssigntraining/Adam/Variable_17training/Adam/zeros_17*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_17*
validate_shape(*
_output_shapes	
:А
Щ
training/Adam/Variable_17/readIdentitytraining/Adam/Variable_17*
T0*,
_class"
 loc:@training/Adam/Variable_17*
_output_shapes	
:А
w
&training/Adam/zeros_18/shape_as_tensorConst*
_output_shapes
:*
valueB"А   (   *
dtype0
a
training/Adam/zeros_18/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
а
training/Adam/zeros_18Fill&training/Adam/zeros_18/shape_as_tensortraining/Adam/zeros_18/Const*
T0*

index_type0*
_output_shapes
:	А(
П
training/Adam/Variable_18
VariableV2*
_output_shapes
:	А(*
	container *
shape:	А(*
shared_name *
dtype0
▐
 training/Adam/Variable_18/AssignAssigntraining/Adam/Variable_18training/Adam/zeros_18*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_18*
validate_shape(*
_output_shapes
:	А(
Э
training/Adam/Variable_18/readIdentitytraining/Adam/Variable_18*,
_class"
 loc:@training/Adam/Variable_18*
_output_shapes
:	А(*
T0
c
training/Adam/zeros_19Const*
valueB(*    *
dtype0*
_output_shapes
:(
Е
training/Adam/Variable_19
VariableV2*
shape:(*
shared_name *
dtype0*
_output_shapes
:(*
	container 
┘
 training/Adam/Variable_19/AssignAssigntraining/Adam/Variable_19training/Adam/zeros_19*
validate_shape(*
_output_shapes
:(*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_19
Ш
training/Adam/Variable_19/readIdentitytraining/Adam/Variable_19*
_output_shapes
:(*
T0*,
_class"
 loc:@training/Adam/Variable_19
p
&training/Adam/zeros_20/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_20/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_20Fill&training/Adam/zeros_20/shape_as_tensortraining/Adam/zeros_20/Const*
T0*

index_type0*
_output_shapes
:
Е
training/Adam/Variable_20
VariableV2*
dtype0*
_output_shapes
:*
	container *
shape:*
shared_name 
┘
 training/Adam/Variable_20/AssignAssigntraining/Adam/Variable_20training/Adam/zeros_20*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_20*
validate_shape(*
_output_shapes
:
Ш
training/Adam/Variable_20/readIdentitytraining/Adam/Variable_20*
T0*,
_class"
 loc:@training/Adam/Variable_20*
_output_shapes
:
p
&training/Adam/zeros_21/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_21/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_21Fill&training/Adam/zeros_21/shape_as_tensortraining/Adam/zeros_21/Const*
T0*

index_type0*
_output_shapes
:
Е
training/Adam/Variable_21
VariableV2*
shape:*
shared_name *
dtype0*
_output_shapes
:*
	container 
┘
 training/Adam/Variable_21/AssignAssigntraining/Adam/Variable_21training/Adam/zeros_21*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_21*
validate_shape(*
_output_shapes
:
Ш
training/Adam/Variable_21/readIdentitytraining/Adam/Variable_21*
T0*,
_class"
 loc:@training/Adam/Variable_21*
_output_shapes
:
p
&training/Adam/zeros_22/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_22/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_22Fill&training/Adam/zeros_22/shape_as_tensortraining/Adam/zeros_22/Const*
T0*

index_type0*
_output_shapes
:
Е
training/Adam/Variable_22
VariableV2*
dtype0*
_output_shapes
:*
	container *
shape:*
shared_name 
┘
 training/Adam/Variable_22/AssignAssigntraining/Adam/Variable_22training/Adam/zeros_22*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_22*
validate_shape(*
_output_shapes
:
Ш
training/Adam/Variable_22/readIdentitytraining/Adam/Variable_22*
_output_shapes
:*
T0*,
_class"
 loc:@training/Adam/Variable_22
p
&training/Adam/zeros_23/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_23/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_23Fill&training/Adam/zeros_23/shape_as_tensortraining/Adam/zeros_23/Const*
_output_shapes
:*
T0*

index_type0
Е
training/Adam/Variable_23
VariableV2*
dtype0*
_output_shapes
:*
	container *
shape:*
shared_name 
┘
 training/Adam/Variable_23/AssignAssigntraining/Adam/Variable_23training/Adam/zeros_23*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_23*
validate_shape(*
_output_shapes
:
Ш
training/Adam/Variable_23/readIdentitytraining/Adam/Variable_23*
T0*,
_class"
 loc:@training/Adam/Variable_23*
_output_shapes
:
p
&training/Adam/zeros_24/shape_as_tensorConst*
_output_shapes
:*
valueB:*
dtype0
a
training/Adam/zeros_24/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_24Fill&training/Adam/zeros_24/shape_as_tensortraining/Adam/zeros_24/Const*
_output_shapes
:*
T0*

index_type0
Е
training/Adam/Variable_24
VariableV2*
shared_name *
dtype0*
_output_shapes
:*
	container *
shape:
┘
 training/Adam/Variable_24/AssignAssigntraining/Adam/Variable_24training/Adam/zeros_24*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_24*
validate_shape(*
_output_shapes
:
Ш
training/Adam/Variable_24/readIdentitytraining/Adam/Variable_24*
T0*,
_class"
 loc:@training/Adam/Variable_24*
_output_shapes
:
p
&training/Adam/zeros_25/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_25/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_25Fill&training/Adam/zeros_25/shape_as_tensortraining/Adam/zeros_25/Const*
_output_shapes
:*
T0*

index_type0
Е
training/Adam/Variable_25
VariableV2*
shape:*
shared_name *
dtype0*
_output_shapes
:*
	container 
┘
 training/Adam/Variable_25/AssignAssigntraining/Adam/Variable_25training/Adam/zeros_25*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_25*
validate_shape(*
_output_shapes
:
Ш
training/Adam/Variable_25/readIdentitytraining/Adam/Variable_25*
T0*,
_class"
 loc:@training/Adam/Variable_25*
_output_shapes
:
p
&training/Adam/zeros_26/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_26/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_26Fill&training/Adam/zeros_26/shape_as_tensortraining/Adam/zeros_26/Const*
T0*

index_type0*
_output_shapes
:
Е
training/Adam/Variable_26
VariableV2*
shape:*
shared_name *
dtype0*
_output_shapes
:*
	container 
┘
 training/Adam/Variable_26/AssignAssigntraining/Adam/Variable_26training/Adam/zeros_26*
validate_shape(*
_output_shapes
:*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_26
Ш
training/Adam/Variable_26/readIdentitytraining/Adam/Variable_26*
_output_shapes
:*
T0*,
_class"
 loc:@training/Adam/Variable_26
p
&training/Adam/zeros_27/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_27/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_27Fill&training/Adam/zeros_27/shape_as_tensortraining/Adam/zeros_27/Const*
_output_shapes
:*
T0*

index_type0
Е
training/Adam/Variable_27
VariableV2*
dtype0*
_output_shapes
:*
	container *
shape:*
shared_name 
┘
 training/Adam/Variable_27/AssignAssigntraining/Adam/Variable_27training/Adam/zeros_27*
T0*,
_class"
 loc:@training/Adam/Variable_27*
validate_shape(*
_output_shapes
:*
use_locking(
Ш
training/Adam/Variable_27/readIdentitytraining/Adam/Variable_27*
T0*,
_class"
 loc:@training/Adam/Variable_27*
_output_shapes
:
p
&training/Adam/zeros_28/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_28/ConstConst*
_output_shapes
: *
valueB
 *    *
dtype0
Ы
training/Adam/zeros_28Fill&training/Adam/zeros_28/shape_as_tensortraining/Adam/zeros_28/Const*
_output_shapes
:*
T0*

index_type0
Е
training/Adam/Variable_28
VariableV2*
shared_name *
dtype0*
_output_shapes
:*
	container *
shape:
┘
 training/Adam/Variable_28/AssignAssigntraining/Adam/Variable_28training/Adam/zeros_28*
validate_shape(*
_output_shapes
:*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_28
Ш
training/Adam/Variable_28/readIdentitytraining/Adam/Variable_28*
_output_shapes
:*
T0*,
_class"
 loc:@training/Adam/Variable_28
p
&training/Adam/zeros_29/shape_as_tensorConst*
valueB:*
dtype0*
_output_shapes
:
a
training/Adam/zeros_29/ConstConst*
valueB
 *    *
dtype0*
_output_shapes
: 
Ы
training/Adam/zeros_29Fill&training/Adam/zeros_29/shape_as_tensortraining/Adam/zeros_29/Const*
T0*

index_type0*
_output_shapes
:
Е
training/Adam/Variable_29
VariableV2*
shared_name *
dtype0*
_output_shapes
:*
	container *
shape:
┘
 training/Adam/Variable_29/AssignAssigntraining/Adam/Variable_29training/Adam/zeros_29*
_output_shapes
:*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_29*
validate_shape(
Ш
training/Adam/Variable_29/readIdentitytraining/Adam/Variable_29*
T0*,
_class"
 loc:@training/Adam/Variable_29*
_output_shapes
:
s
training/Adam/mul_1MulAdam/beta_1/readtraining/Adam/Variable/read*
_output_shapes
:	zА*
T0
Z
training/Adam/sub_2/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
d
training/Adam/sub_2Subtraining/Adam/sub_2/xAdam/beta_1/read*
_output_shapes
: *
T0
Ч
training/Adam/mul_2Multraining/Adam/sub_2<training/Adam/gradients/input_layer/transpose_grad/transpose*
_output_shapes
:	zА*
T0
n
training/Adam/add_1Addtraining/Adam/mul_1training/Adam/mul_2*
T0*
_output_shapes
:	zА
v
training/Adam/mul_3MulAdam/beta_2/readtraining/Adam/Variable_10/read*
T0*
_output_shapes
:	zА
Z
training/Adam/sub_3/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
d
training/Adam/sub_3Subtraining/Adam/sub_3/xAdam/beta_2/read*
T0*
_output_shapes
: 
Ж
training/Adam/SquareSquare<training/Adam/gradients/input_layer/transpose_grad/transpose*
T0*
_output_shapes
:	zА
o
training/Adam/mul_4Multraining/Adam/sub_3training/Adam/Square*
T0*
_output_shapes
:	zА
n
training/Adam/add_2Addtraining/Adam/mul_3training/Adam/mul_4*
_output_shapes
:	zА*
T0
l
training/Adam/mul_5Multraining/Adam/multraining/Adam/add_1*
T0*
_output_shapes
:	zА
Z
training/Adam/Const_2Const*
_output_shapes
: *
valueB
 *    *
dtype0
Z
training/Adam/Const_3Const*
valueB
 *  А*
dtype0*
_output_shapes
: 
Ж
%training/Adam/clip_by_value_1/MinimumMinimumtraining/Adam/add_2training/Adam/Const_3*
_output_shapes
:	zА*
T0
Р
training/Adam/clip_by_value_1Maximum%training/Adam/clip_by_value_1/Minimumtraining/Adam/Const_2*
_output_shapes
:	zА*
T0
e
training/Adam/Sqrt_1Sqrttraining/Adam/clip_by_value_1*
T0*
_output_shapes
:	zА
Z
training/Adam/add_3/yConst*
_output_shapes
: *
valueB
 *Х┐╓3*
dtype0
q
training/Adam/add_3Addtraining/Adam/Sqrt_1training/Adam/add_3/y*
T0*
_output_shapes
:	zА
v
training/Adam/truediv_1RealDivtraining/Adam/mul_5training/Adam/add_3*
_output_shapes
:	zА*
T0
v
training/Adam/sub_4Subinput_layer/kernel/readtraining/Adam/truediv_1*
T0*
_output_shapes
:	zА
╔
training/Adam/AssignAssigntraining/Adam/Variabletraining/Adam/add_1*)
_class
loc:@training/Adam/Variable*
validate_shape(*
_output_shapes
:	zА*
use_locking(*
T0
╤
training/Adam/Assign_1Assigntraining/Adam/Variable_10training/Adam/add_2*
_output_shapes
:	zА*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_10*
validate_shape(
├
training/Adam/Assign_2Assigninput_layer/kerneltraining/Adam/sub_4*
use_locking(*
T0*%
_class
loc:@input_layer/kernel*
validate_shape(*
_output_shapes
:	zА
q
training/Adam/mul_6MulAdam/beta_1/readtraining/Adam/Variable_1/read*
T0*
_output_shapes	
:А
Z
training/Adam/sub_5/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
d
training/Adam/sub_5Subtraining/Adam/sub_5/xAdam/beta_1/read*
_output_shapes
: *
T0
С
training/Adam/mul_7Multraining/Adam/sub_5:training/Adam/gradients/input_layer/Reshape_3_grad/Reshape*
T0*
_output_shapes	
:А
j
training/Adam/add_4Addtraining/Adam/mul_6training/Adam/mul_7*
T0*
_output_shapes	
:А
r
training/Adam/mul_8MulAdam/beta_2/readtraining/Adam/Variable_11/read*
T0*
_output_shapes	
:А
Z
training/Adam/sub_6/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
d
training/Adam/sub_6Subtraining/Adam/sub_6/xAdam/beta_2/read*
_output_shapes
: *
T0
В
training/Adam/Square_1Square:training/Adam/gradients/input_layer/Reshape_3_grad/Reshape*
_output_shapes	
:А*
T0
m
training/Adam/mul_9Multraining/Adam/sub_6training/Adam/Square_1*
T0*
_output_shapes	
:А
j
training/Adam/add_5Addtraining/Adam/mul_8training/Adam/mul_9*
_output_shapes	
:А*
T0
i
training/Adam/mul_10Multraining/Adam/multraining/Adam/add_4*
_output_shapes	
:А*
T0
Z
training/Adam/Const_4Const*
valueB
 *    *
dtype0*
_output_shapes
: 
Z
training/Adam/Const_5Const*
valueB
 *  А*
dtype0*
_output_shapes
: 
В
%training/Adam/clip_by_value_2/MinimumMinimumtraining/Adam/add_5training/Adam/Const_5*
T0*
_output_shapes	
:А
М
training/Adam/clip_by_value_2Maximum%training/Adam/clip_by_value_2/Minimumtraining/Adam/Const_4*
T0*
_output_shapes	
:А
a
training/Adam/Sqrt_2Sqrttraining/Adam/clip_by_value_2*
T0*
_output_shapes	
:А
Z
training/Adam/add_6/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
m
training/Adam/add_6Addtraining/Adam/Sqrt_2training/Adam/add_6/y*
T0*
_output_shapes	
:А
s
training/Adam/truediv_2RealDivtraining/Adam/mul_10training/Adam/add_6*
T0*
_output_shapes	
:А
p
training/Adam/sub_7Subinput_layer/bias/readtraining/Adam/truediv_2*
T0*
_output_shapes	
:А
╦
training/Adam/Assign_3Assigntraining/Adam/Variable_1training/Adam/add_4*
T0*+
_class!
loc:@training/Adam/Variable_1*
validate_shape(*
_output_shapes	
:А*
use_locking(
═
training/Adam/Assign_4Assigntraining/Adam/Variable_11training/Adam/add_5*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_11*
validate_shape(*
_output_shapes	
:А
╗
training/Adam/Assign_5Assigninput_layer/biastraining/Adam/sub_7*
use_locking(*
T0*#
_class
loc:@input_layer/bias*
validate_shape(*
_output_shapes	
:А
w
training/Adam/mul_11MulAdam/beta_1/readtraining/Adam/Variable_2/read*
T0* 
_output_shapes
:
АА
Z
training/Adam/sub_8/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
d
training/Adam/sub_8Subtraining/Adam/sub_8/xAdam/beta_1/read*
T0*
_output_shapes
: 
Х
training/Adam/mul_12Multraining/Adam/sub_88training/Adam/gradients/dense_1/transpose_grad/transpose*
T0* 
_output_shapes
:
АА
q
training/Adam/add_7Addtraining/Adam/mul_11training/Adam/mul_12*
T0* 
_output_shapes
:
АА
x
training/Adam/mul_13MulAdam/beta_2/readtraining/Adam/Variable_12/read*
T0* 
_output_shapes
:
АА
Z
training/Adam/sub_9/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
d
training/Adam/sub_9Subtraining/Adam/sub_9/xAdam/beta_2/read*
T0*
_output_shapes
: 
Е
training/Adam/Square_2Square8training/Adam/gradients/dense_1/transpose_grad/transpose*
T0* 
_output_shapes
:
АА
s
training/Adam/mul_14Multraining/Adam/sub_9training/Adam/Square_2*
T0* 
_output_shapes
:
АА
q
training/Adam/add_8Addtraining/Adam/mul_13training/Adam/mul_14* 
_output_shapes
:
АА*
T0
n
training/Adam/mul_15Multraining/Adam/multraining/Adam/add_7*
T0* 
_output_shapes
:
АА
Z
training/Adam/Const_6Const*
valueB
 *    *
dtype0*
_output_shapes
: 
Z
training/Adam/Const_7Const*
dtype0*
_output_shapes
: *
valueB
 *  А
З
%training/Adam/clip_by_value_3/MinimumMinimumtraining/Adam/add_8training/Adam/Const_7*
T0* 
_output_shapes
:
АА
С
training/Adam/clip_by_value_3Maximum%training/Adam/clip_by_value_3/Minimumtraining/Adam/Const_6*
T0* 
_output_shapes
:
АА
f
training/Adam/Sqrt_3Sqrttraining/Adam/clip_by_value_3* 
_output_shapes
:
АА*
T0
Z
training/Adam/add_9/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
r
training/Adam/add_9Addtraining/Adam/Sqrt_3training/Adam/add_9/y*
T0* 
_output_shapes
:
АА
x
training/Adam/truediv_3RealDivtraining/Adam/mul_15training/Adam/add_9* 
_output_shapes
:
АА*
T0
t
training/Adam/sub_10Subdense_1/kernel/readtraining/Adam/truediv_3*
T0* 
_output_shapes
:
АА
╨
training/Adam/Assign_6Assigntraining/Adam/Variable_2training/Adam/add_7*+
_class!
loc:@training/Adam/Variable_2*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0
╥
training/Adam/Assign_7Assigntraining/Adam/Variable_12training/Adam/add_8* 
_output_shapes
:
АА*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_12*
validate_shape(
╜
training/Adam/Assign_8Assigndense_1/kerneltraining/Adam/sub_10*
use_locking(*
T0*!
_class
loc:@dense_1/kernel*
validate_shape(* 
_output_shapes
:
АА
r
training/Adam/mul_16MulAdam/beta_1/readtraining/Adam/Variable_3/read*
_output_shapes	
:А*
T0
[
training/Adam/sub_11/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_11Subtraining/Adam/sub_11/xAdam/beta_1/read*
T0*
_output_shapes
: 
П
training/Adam/mul_17Multraining/Adam/sub_116training/Adam/gradients/dense_1/Reshape_3_grad/Reshape*
T0*
_output_shapes	
:А
m
training/Adam/add_10Addtraining/Adam/mul_16training/Adam/mul_17*
T0*
_output_shapes	
:А
s
training/Adam/mul_18MulAdam/beta_2/readtraining/Adam/Variable_13/read*
_output_shapes	
:А*
T0
[
training/Adam/sub_12/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_12Subtraining/Adam/sub_12/xAdam/beta_2/read*
_output_shapes
: *
T0
~
training/Adam/Square_3Square6training/Adam/gradients/dense_1/Reshape_3_grad/Reshape*
T0*
_output_shapes	
:А
o
training/Adam/mul_19Multraining/Adam/sub_12training/Adam/Square_3*
T0*
_output_shapes	
:А
m
training/Adam/add_11Addtraining/Adam/mul_18training/Adam/mul_19*
T0*
_output_shapes	
:А
j
training/Adam/mul_20Multraining/Adam/multraining/Adam/add_10*
T0*
_output_shapes	
:А
Z
training/Adam/Const_8Const*
valueB
 *    *
dtype0*
_output_shapes
: 
Z
training/Adam/Const_9Const*
dtype0*
_output_shapes
: *
valueB
 *  А
Г
%training/Adam/clip_by_value_4/MinimumMinimumtraining/Adam/add_11training/Adam/Const_9*
T0*
_output_shapes	
:А
М
training/Adam/clip_by_value_4Maximum%training/Adam/clip_by_value_4/Minimumtraining/Adam/Const_8*
T0*
_output_shapes	
:А
a
training/Adam/Sqrt_4Sqrttraining/Adam/clip_by_value_4*
T0*
_output_shapes	
:А
[
training/Adam/add_12/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
o
training/Adam/add_12Addtraining/Adam/Sqrt_4training/Adam/add_12/y*
_output_shapes	
:А*
T0
t
training/Adam/truediv_4RealDivtraining/Adam/mul_20training/Adam/add_12*
T0*
_output_shapes	
:А
m
training/Adam/sub_13Subdense_1/bias/readtraining/Adam/truediv_4*
T0*
_output_shapes	
:А
╠
training/Adam/Assign_9Assigntraining/Adam/Variable_3training/Adam/add_10*
T0*+
_class!
loc:@training/Adam/Variable_3*
validate_shape(*
_output_shapes	
:А*
use_locking(
╧
training/Adam/Assign_10Assigntraining/Adam/Variable_13training/Adam/add_11*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_13*
validate_shape(*
_output_shapes	
:А
╡
training/Adam/Assign_11Assigndense_1/biastraining/Adam/sub_13*
use_locking(*
T0*
_class
loc:@dense_1/bias*
validate_shape(*
_output_shapes	
:А
w
training/Adam/mul_21MulAdam/beta_1/readtraining/Adam/Variable_4/read*
T0* 
_output_shapes
:
АА
[
training/Adam/sub_14/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_14Subtraining/Adam/sub_14/xAdam/beta_1/read*
_output_shapes
: *
T0
Ц
training/Adam/mul_22Multraining/Adam/sub_148training/Adam/gradients/dense_2/transpose_grad/transpose* 
_output_shapes
:
АА*
T0
r
training/Adam/add_13Addtraining/Adam/mul_21training/Adam/mul_22*
T0* 
_output_shapes
:
АА
x
training/Adam/mul_23MulAdam/beta_2/readtraining/Adam/Variable_14/read*
T0* 
_output_shapes
:
АА
[
training/Adam/sub_15/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_15Subtraining/Adam/sub_15/xAdam/beta_2/read*
T0*
_output_shapes
: 
Е
training/Adam/Square_4Square8training/Adam/gradients/dense_2/transpose_grad/transpose*
T0* 
_output_shapes
:
АА
t
training/Adam/mul_24Multraining/Adam/sub_15training/Adam/Square_4* 
_output_shapes
:
АА*
T0
r
training/Adam/add_14Addtraining/Adam/mul_23training/Adam/mul_24*
T0* 
_output_shapes
:
АА
o
training/Adam/mul_25Multraining/Adam/multraining/Adam/add_13*
T0* 
_output_shapes
:
АА
[
training/Adam/Const_10Const*
valueB
 *    *
dtype0*
_output_shapes
: 
[
training/Adam/Const_11Const*
valueB
 *  А*
dtype0*
_output_shapes
: 
Й
%training/Adam/clip_by_value_5/MinimumMinimumtraining/Adam/add_14training/Adam/Const_11* 
_output_shapes
:
АА*
T0
Т
training/Adam/clip_by_value_5Maximum%training/Adam/clip_by_value_5/Minimumtraining/Adam/Const_10*
T0* 
_output_shapes
:
АА
f
training/Adam/Sqrt_5Sqrttraining/Adam/clip_by_value_5* 
_output_shapes
:
АА*
T0
[
training/Adam/add_15/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
t
training/Adam/add_15Addtraining/Adam/Sqrt_5training/Adam/add_15/y*
T0* 
_output_shapes
:
АА
y
training/Adam/truediv_5RealDivtraining/Adam/mul_25training/Adam/add_15*
T0* 
_output_shapes
:
АА
t
training/Adam/sub_16Subdense_2/kernel/readtraining/Adam/truediv_5*
T0* 
_output_shapes
:
АА
╥
training/Adam/Assign_12Assigntraining/Adam/Variable_4training/Adam/add_13*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_4*
validate_shape(* 
_output_shapes
:
АА
╘
training/Adam/Assign_13Assigntraining/Adam/Variable_14training/Adam/add_14*
T0*,
_class"
 loc:@training/Adam/Variable_14*
validate_shape(* 
_output_shapes
:
АА*
use_locking(
╛
training/Adam/Assign_14Assigndense_2/kerneltraining/Adam/sub_16*
use_locking(*
T0*!
_class
loc:@dense_2/kernel*
validate_shape(* 
_output_shapes
:
АА
r
training/Adam/mul_26MulAdam/beta_1/readtraining/Adam/Variable_5/read*
T0*
_output_shapes	
:А
[
training/Adam/sub_17/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_17Subtraining/Adam/sub_17/xAdam/beta_1/read*
T0*
_output_shapes
: 
П
training/Adam/mul_27Multraining/Adam/sub_176training/Adam/gradients/dense_2/Reshape_3_grad/Reshape*
T0*
_output_shapes	
:А
m
training/Adam/add_16Addtraining/Adam/mul_26training/Adam/mul_27*
T0*
_output_shapes	
:А
s
training/Adam/mul_28MulAdam/beta_2/readtraining/Adam/Variable_15/read*
T0*
_output_shapes	
:А
[
training/Adam/sub_18/xConst*
_output_shapes
: *
valueB
 *  А?*
dtype0
f
training/Adam/sub_18Subtraining/Adam/sub_18/xAdam/beta_2/read*
_output_shapes
: *
T0
~
training/Adam/Square_5Square6training/Adam/gradients/dense_2/Reshape_3_grad/Reshape*
T0*
_output_shapes	
:А
o
training/Adam/mul_29Multraining/Adam/sub_18training/Adam/Square_5*
T0*
_output_shapes	
:А
m
training/Adam/add_17Addtraining/Adam/mul_28training/Adam/mul_29*
T0*
_output_shapes	
:А
j
training/Adam/mul_30Multraining/Adam/multraining/Adam/add_16*
T0*
_output_shapes	
:А
[
training/Adam/Const_12Const*
valueB
 *    *
dtype0*
_output_shapes
: 
[
training/Adam/Const_13Const*
valueB
 *  А*
dtype0*
_output_shapes
: 
Д
%training/Adam/clip_by_value_6/MinimumMinimumtraining/Adam/add_17training/Adam/Const_13*
_output_shapes	
:А*
T0
Н
training/Adam/clip_by_value_6Maximum%training/Adam/clip_by_value_6/Minimumtraining/Adam/Const_12*
T0*
_output_shapes	
:А
a
training/Adam/Sqrt_6Sqrttraining/Adam/clip_by_value_6*
_output_shapes	
:А*
T0
[
training/Adam/add_18/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
o
training/Adam/add_18Addtraining/Adam/Sqrt_6training/Adam/add_18/y*
_output_shapes	
:А*
T0
t
training/Adam/truediv_6RealDivtraining/Adam/mul_30training/Adam/add_18*
T0*
_output_shapes	
:А
m
training/Adam/sub_19Subdense_2/bias/readtraining/Adam/truediv_6*
_output_shapes	
:А*
T0
═
training/Adam/Assign_15Assigntraining/Adam/Variable_5training/Adam/add_16*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_5*
validate_shape(*
_output_shapes	
:А
╧
training/Adam/Assign_16Assigntraining/Adam/Variable_15training/Adam/add_17*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_15*
validate_shape(*
_output_shapes	
:А
╡
training/Adam/Assign_17Assigndense_2/biastraining/Adam/sub_19*
_output_shapes	
:А*
use_locking(*
T0*
_class
loc:@dense_2/bias*
validate_shape(
w
training/Adam/mul_31MulAdam/beta_1/readtraining/Adam/Variable_6/read*
T0* 
_output_shapes
:
АА
[
training/Adam/sub_20/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_20Subtraining/Adam/sub_20/xAdam/beta_1/read*
_output_shapes
: *
T0
Ц
training/Adam/mul_32Multraining/Adam/sub_208training/Adam/gradients/dense_3/transpose_grad/transpose*
T0* 
_output_shapes
:
АА
r
training/Adam/add_19Addtraining/Adam/mul_31training/Adam/mul_32* 
_output_shapes
:
АА*
T0
x
training/Adam/mul_33MulAdam/beta_2/readtraining/Adam/Variable_16/read* 
_output_shapes
:
АА*
T0
[
training/Adam/sub_21/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_21Subtraining/Adam/sub_21/xAdam/beta_2/read*
_output_shapes
: *
T0
Е
training/Adam/Square_6Square8training/Adam/gradients/dense_3/transpose_grad/transpose* 
_output_shapes
:
АА*
T0
t
training/Adam/mul_34Multraining/Adam/sub_21training/Adam/Square_6* 
_output_shapes
:
АА*
T0
r
training/Adam/add_20Addtraining/Adam/mul_33training/Adam/mul_34*
T0* 
_output_shapes
:
АА
o
training/Adam/mul_35Multraining/Adam/multraining/Adam/add_19*
T0* 
_output_shapes
:
АА
[
training/Adam/Const_14Const*
valueB
 *    *
dtype0*
_output_shapes
: 
[
training/Adam/Const_15Const*
_output_shapes
: *
valueB
 *  А*
dtype0
Й
%training/Adam/clip_by_value_7/MinimumMinimumtraining/Adam/add_20training/Adam/Const_15*
T0* 
_output_shapes
:
АА
Т
training/Adam/clip_by_value_7Maximum%training/Adam/clip_by_value_7/Minimumtraining/Adam/Const_14* 
_output_shapes
:
АА*
T0
f
training/Adam/Sqrt_7Sqrttraining/Adam/clip_by_value_7*
T0* 
_output_shapes
:
АА
[
training/Adam/add_21/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
t
training/Adam/add_21Addtraining/Adam/Sqrt_7training/Adam/add_21/y*
T0* 
_output_shapes
:
АА
y
training/Adam/truediv_7RealDivtraining/Adam/mul_35training/Adam/add_21*
T0* 
_output_shapes
:
АА
t
training/Adam/sub_22Subdense_3/kernel/readtraining/Adam/truediv_7*
T0* 
_output_shapes
:
АА
╥
training/Adam/Assign_18Assigntraining/Adam/Variable_6training/Adam/add_19*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_6*
validate_shape(* 
_output_shapes
:
АА
╘
training/Adam/Assign_19Assigntraining/Adam/Variable_16training/Adam/add_20*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_16
╛
training/Adam/Assign_20Assigndense_3/kerneltraining/Adam/sub_22*!
_class
loc:@dense_3/kernel*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0
r
training/Adam/mul_36MulAdam/beta_1/readtraining/Adam/Variable_7/read*
_output_shapes	
:А*
T0
[
training/Adam/sub_23/xConst*
dtype0*
_output_shapes
: *
valueB
 *  А?
f
training/Adam/sub_23Subtraining/Adam/sub_23/xAdam/beta_1/read*
T0*
_output_shapes
: 
П
training/Adam/mul_37Multraining/Adam/sub_236training/Adam/gradients/dense_3/Reshape_3_grad/Reshape*
T0*
_output_shapes	
:А
m
training/Adam/add_22Addtraining/Adam/mul_36training/Adam/mul_37*
T0*
_output_shapes	
:А
s
training/Adam/mul_38MulAdam/beta_2/readtraining/Adam/Variable_17/read*
T0*
_output_shapes	
:А
[
training/Adam/sub_24/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_24Subtraining/Adam/sub_24/xAdam/beta_2/read*
_output_shapes
: *
T0
~
training/Adam/Square_7Square6training/Adam/gradients/dense_3/Reshape_3_grad/Reshape*
T0*
_output_shapes	
:А
o
training/Adam/mul_39Multraining/Adam/sub_24training/Adam/Square_7*
_output_shapes	
:А*
T0
m
training/Adam/add_23Addtraining/Adam/mul_38training/Adam/mul_39*
_output_shapes	
:А*
T0
j
training/Adam/mul_40Multraining/Adam/multraining/Adam/add_22*
T0*
_output_shapes	
:А
[
training/Adam/Const_16Const*
valueB
 *    *
dtype0*
_output_shapes
: 
[
training/Adam/Const_17Const*
_output_shapes
: *
valueB
 *  А*
dtype0
Д
%training/Adam/clip_by_value_8/MinimumMinimumtraining/Adam/add_23training/Adam/Const_17*
T0*
_output_shapes	
:А
Н
training/Adam/clip_by_value_8Maximum%training/Adam/clip_by_value_8/Minimumtraining/Adam/Const_16*
_output_shapes	
:А*
T0
a
training/Adam/Sqrt_8Sqrttraining/Adam/clip_by_value_8*
_output_shapes	
:А*
T0
[
training/Adam/add_24/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
o
training/Adam/add_24Addtraining/Adam/Sqrt_8training/Adam/add_24/y*
_output_shapes	
:А*
T0
t
training/Adam/truediv_8RealDivtraining/Adam/mul_40training/Adam/add_24*
_output_shapes	
:А*
T0
m
training/Adam/sub_25Subdense_3/bias/readtraining/Adam/truediv_8*
_output_shapes	
:А*
T0
═
training/Adam/Assign_21Assigntraining/Adam/Variable_7training/Adam/add_22*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_7*
validate_shape(*
_output_shapes	
:А
╧
training/Adam/Assign_22Assigntraining/Adam/Variable_17training/Adam/add_23*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_17*
validate_shape(*
_output_shapes	
:А
╡
training/Adam/Assign_23Assigndense_3/biastraining/Adam/sub_25*
use_locking(*
T0*
_class
loc:@dense_3/bias*
validate_shape(*
_output_shapes	
:А
v
training/Adam/mul_41MulAdam/beta_1/readtraining/Adam/Variable_8/read*
_output_shapes
:	А(*
T0
[
training/Adam/sub_26/xConst*
_output_shapes
: *
valueB
 *  А?*
dtype0
f
training/Adam/sub_26Subtraining/Adam/sub_26/xAdam/beta_1/read*
T0*
_output_shapes
: 
Ъ
training/Adam/mul_42Multraining/Adam/sub_26=training/Adam/gradients/output_layer/transpose_grad/transpose*
_output_shapes
:	А(*
T0
q
training/Adam/add_25Addtraining/Adam/mul_41training/Adam/mul_42*
T0*
_output_shapes
:	А(
w
training/Adam/mul_43MulAdam/beta_2/readtraining/Adam/Variable_18/read*
T0*
_output_shapes
:	А(
[
training/Adam/sub_27/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_27Subtraining/Adam/sub_27/xAdam/beta_2/read*
_output_shapes
: *
T0
Й
training/Adam/Square_8Square=training/Adam/gradients/output_layer/transpose_grad/transpose*
T0*
_output_shapes
:	А(
s
training/Adam/mul_44Multraining/Adam/sub_27training/Adam/Square_8*
_output_shapes
:	А(*
T0
q
training/Adam/add_26Addtraining/Adam/mul_43training/Adam/mul_44*
T0*
_output_shapes
:	А(
n
training/Adam/mul_45Multraining/Adam/multraining/Adam/add_25*
T0*
_output_shapes
:	А(
[
training/Adam/Const_18Const*
valueB
 *    *
dtype0*
_output_shapes
: 
[
training/Adam/Const_19Const*
valueB
 *  А*
dtype0*
_output_shapes
: 
И
%training/Adam/clip_by_value_9/MinimumMinimumtraining/Adam/add_26training/Adam/Const_19*
T0*
_output_shapes
:	А(
С
training/Adam/clip_by_value_9Maximum%training/Adam/clip_by_value_9/Minimumtraining/Adam/Const_18*
T0*
_output_shapes
:	А(
e
training/Adam/Sqrt_9Sqrttraining/Adam/clip_by_value_9*
_output_shapes
:	А(*
T0
[
training/Adam/add_27/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
s
training/Adam/add_27Addtraining/Adam/Sqrt_9training/Adam/add_27/y*
T0*
_output_shapes
:	А(
x
training/Adam/truediv_9RealDivtraining/Adam/mul_45training/Adam/add_27*
_output_shapes
:	А(*
T0
x
training/Adam/sub_28Suboutput_layer/kernel/readtraining/Adam/truediv_9*
T0*
_output_shapes
:	А(
╤
training/Adam/Assign_24Assigntraining/Adam/Variable_8training/Adam/add_25*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_8*
validate_shape(*
_output_shapes
:	А(
╙
training/Adam/Assign_25Assigntraining/Adam/Variable_18training/Adam/add_26*
validate_shape(*
_output_shapes
:	А(*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_18
╟
training/Adam/Assign_26Assignoutput_layer/kerneltraining/Adam/sub_28*
use_locking(*
T0*&
_class
loc:@output_layer/kernel*
validate_shape(*
_output_shapes
:	А(
q
training/Adam/mul_46MulAdam/beta_1/readtraining/Adam/Variable_9/read*
T0*
_output_shapes
:(
[
training/Adam/sub_29/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_29Subtraining/Adam/sub_29/xAdam/beta_1/read*
_output_shapes
: *
T0
У
training/Adam/mul_47Multraining/Adam/sub_29;training/Adam/gradients/output_layer/Reshape_3_grad/Reshape*
T0*
_output_shapes
:(
l
training/Adam/add_28Addtraining/Adam/mul_46training/Adam/mul_47*
T0*
_output_shapes
:(
r
training/Adam/mul_48MulAdam/beta_2/readtraining/Adam/Variable_19/read*
T0*
_output_shapes
:(
[
training/Adam/sub_30/xConst*
valueB
 *  А?*
dtype0*
_output_shapes
: 
f
training/Adam/sub_30Subtraining/Adam/sub_30/xAdam/beta_2/read*
T0*
_output_shapes
: 
В
training/Adam/Square_9Square;training/Adam/gradients/output_layer/Reshape_3_grad/Reshape*
T0*
_output_shapes
:(
n
training/Adam/mul_49Multraining/Adam/sub_30training/Adam/Square_9*
T0*
_output_shapes
:(
l
training/Adam/add_29Addtraining/Adam/mul_48training/Adam/mul_49*
T0*
_output_shapes
:(
i
training/Adam/mul_50Multraining/Adam/multraining/Adam/add_28*
T0*
_output_shapes
:(
[
training/Adam/Const_20Const*
valueB
 *    *
dtype0*
_output_shapes
: 
[
training/Adam/Const_21Const*
valueB
 *  А*
dtype0*
_output_shapes
: 
Д
&training/Adam/clip_by_value_10/MinimumMinimumtraining/Adam/add_29training/Adam/Const_21*
T0*
_output_shapes
:(
О
training/Adam/clip_by_value_10Maximum&training/Adam/clip_by_value_10/Minimumtraining/Adam/Const_20*
_output_shapes
:(*
T0
b
training/Adam/Sqrt_10Sqrttraining/Adam/clip_by_value_10*
T0*
_output_shapes
:(
[
training/Adam/add_30/yConst*
valueB
 *Х┐╓3*
dtype0*
_output_shapes
: 
o
training/Adam/add_30Addtraining/Adam/Sqrt_10training/Adam/add_30/y*
T0*
_output_shapes
:(
t
training/Adam/truediv_10RealDivtraining/Adam/mul_50training/Adam/add_30*
T0*
_output_shapes
:(
r
training/Adam/sub_31Suboutput_layer/bias/readtraining/Adam/truediv_10*
_output_shapes
:(*
T0
╠
training/Adam/Assign_27Assigntraining/Adam/Variable_9training/Adam/add_28*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_9*
validate_shape(*
_output_shapes
:(
╬
training/Adam/Assign_28Assigntraining/Adam/Variable_19training/Adam/add_29*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_19*
validate_shape(*
_output_shapes
:(
╛
training/Adam/Assign_29Assignoutput_layer/biastraining/Adam/sub_31*
T0*$
_class
loc:@output_layer/bias*
validate_shape(*
_output_shapes
:(*
use_locking(
х
training/group_depsNoOp	^loss/mul#^metrics/mean_absolute_error/Mean_1^training/Adam/Assign^training/Adam/AssignAdd^training/Adam/Assign_1^training/Adam/Assign_10^training/Adam/Assign_11^training/Adam/Assign_12^training/Adam/Assign_13^training/Adam/Assign_14^training/Adam/Assign_15^training/Adam/Assign_16^training/Adam/Assign_17^training/Adam/Assign_18^training/Adam/Assign_19^training/Adam/Assign_2^training/Adam/Assign_20^training/Adam/Assign_21^training/Adam/Assign_22^training/Adam/Assign_23^training/Adam/Assign_24^training/Adam/Assign_25^training/Adam/Assign_26^training/Adam/Assign_27^training/Adam/Assign_28^training/Adam/Assign_29^training/Adam/Assign_3^training/Adam/Assign_4^training/Adam/Assign_5^training/Adam/Assign_6^training/Adam/Assign_7^training/Adam/Assign_8^training/Adam/Assign_9
B

group_depsNoOp	^loss/mul#^metrics/mean_absolute_error/Mean_1
О
IsVariableInitializedIsVariableInitializedinput_layer/kernel*%
_class
loc:@input_layer/kernel*
dtype0*
_output_shapes
: 
М
IsVariableInitialized_1IsVariableInitializedinput_layer/bias*#
_class
loc:@input_layer/bias*
dtype0*
_output_shapes
: 
И
IsVariableInitialized_2IsVariableInitializeddense_1/kernel*
_output_shapes
: *!
_class
loc:@dense_1/kernel*
dtype0
Д
IsVariableInitialized_3IsVariableInitializeddense_1/bias*
_output_shapes
: *
_class
loc:@dense_1/bias*
dtype0
И
IsVariableInitialized_4IsVariableInitializeddense_2/kernel*
_output_shapes
: *!
_class
loc:@dense_2/kernel*
dtype0
Д
IsVariableInitialized_5IsVariableInitializeddense_2/bias*
_class
loc:@dense_2/bias*
dtype0*
_output_shapes
: 
И
IsVariableInitialized_6IsVariableInitializeddense_3/kernel*!
_class
loc:@dense_3/kernel*
dtype0*
_output_shapes
: 
Д
IsVariableInitialized_7IsVariableInitializeddense_3/bias*
_class
loc:@dense_3/bias*
dtype0*
_output_shapes
: 
Т
IsVariableInitialized_8IsVariableInitializedoutput_layer/kernel*
_output_shapes
: *&
_class
loc:@output_layer/kernel*
dtype0
О
IsVariableInitialized_9IsVariableInitializedoutput_layer/bias*$
_class
loc:@output_layer/bias*
dtype0*
_output_shapes
: 
Л
IsVariableInitialized_10IsVariableInitializedAdam/iterations*"
_class
loc:@Adam/iterations*
dtype0	*
_output_shapes
: 
{
IsVariableInitialized_11IsVariableInitializedAdam/lr*
_class
loc:@Adam/lr*
dtype0*
_output_shapes
: 
Г
IsVariableInitialized_12IsVariableInitializedAdam/beta_1*
_class
loc:@Adam/beta_1*
dtype0*
_output_shapes
: 
Г
IsVariableInitialized_13IsVariableInitializedAdam/beta_2*
_class
loc:@Adam/beta_2*
dtype0*
_output_shapes
: 
Б
IsVariableInitialized_14IsVariableInitialized
Adam/decay*
_output_shapes
: *
_class
loc:@Adam/decay*
dtype0
Щ
IsVariableInitialized_15IsVariableInitializedtraining/Adam/Variable*
_output_shapes
: *)
_class
loc:@training/Adam/Variable*
dtype0
Э
IsVariableInitialized_16IsVariableInitializedtraining/Adam/Variable_1*
dtype0*
_output_shapes
: *+
_class!
loc:@training/Adam/Variable_1
Э
IsVariableInitialized_17IsVariableInitializedtraining/Adam/Variable_2*+
_class!
loc:@training/Adam/Variable_2*
dtype0*
_output_shapes
: 
Э
IsVariableInitialized_18IsVariableInitializedtraining/Adam/Variable_3*+
_class!
loc:@training/Adam/Variable_3*
dtype0*
_output_shapes
: 
Э
IsVariableInitialized_19IsVariableInitializedtraining/Adam/Variable_4*+
_class!
loc:@training/Adam/Variable_4*
dtype0*
_output_shapes
: 
Э
IsVariableInitialized_20IsVariableInitializedtraining/Adam/Variable_5*+
_class!
loc:@training/Adam/Variable_5*
dtype0*
_output_shapes
: 
Э
IsVariableInitialized_21IsVariableInitializedtraining/Adam/Variable_6*+
_class!
loc:@training/Adam/Variable_6*
dtype0*
_output_shapes
: 
Э
IsVariableInitialized_22IsVariableInitializedtraining/Adam/Variable_7*+
_class!
loc:@training/Adam/Variable_7*
dtype0*
_output_shapes
: 
Э
IsVariableInitialized_23IsVariableInitializedtraining/Adam/Variable_8*
_output_shapes
: *+
_class!
loc:@training/Adam/Variable_8*
dtype0
Э
IsVariableInitialized_24IsVariableInitializedtraining/Adam/Variable_9*+
_class!
loc:@training/Adam/Variable_9*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_25IsVariableInitializedtraining/Adam/Variable_10*,
_class"
 loc:@training/Adam/Variable_10*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_26IsVariableInitializedtraining/Adam/Variable_11*,
_class"
 loc:@training/Adam/Variable_11*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_27IsVariableInitializedtraining/Adam/Variable_12*,
_class"
 loc:@training/Adam/Variable_12*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_28IsVariableInitializedtraining/Adam/Variable_13*
dtype0*
_output_shapes
: *,
_class"
 loc:@training/Adam/Variable_13
Я
IsVariableInitialized_29IsVariableInitializedtraining/Adam/Variable_14*,
_class"
 loc:@training/Adam/Variable_14*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_30IsVariableInitializedtraining/Adam/Variable_15*,
_class"
 loc:@training/Adam/Variable_15*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_31IsVariableInitializedtraining/Adam/Variable_16*,
_class"
 loc:@training/Adam/Variable_16*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_32IsVariableInitializedtraining/Adam/Variable_17*
dtype0*
_output_shapes
: *,
_class"
 loc:@training/Adam/Variable_17
Я
IsVariableInitialized_33IsVariableInitializedtraining/Adam/Variable_18*,
_class"
 loc:@training/Adam/Variable_18*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_34IsVariableInitializedtraining/Adam/Variable_19*,
_class"
 loc:@training/Adam/Variable_19*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_35IsVariableInitializedtraining/Adam/Variable_20*,
_class"
 loc:@training/Adam/Variable_20*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_36IsVariableInitializedtraining/Adam/Variable_21*,
_class"
 loc:@training/Adam/Variable_21*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_37IsVariableInitializedtraining/Adam/Variable_22*,
_class"
 loc:@training/Adam/Variable_22*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_38IsVariableInitializedtraining/Adam/Variable_23*,
_class"
 loc:@training/Adam/Variable_23*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_39IsVariableInitializedtraining/Adam/Variable_24*,
_class"
 loc:@training/Adam/Variable_24*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_40IsVariableInitializedtraining/Adam/Variable_25*,
_class"
 loc:@training/Adam/Variable_25*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_41IsVariableInitializedtraining/Adam/Variable_26*
dtype0*
_output_shapes
: *,
_class"
 loc:@training/Adam/Variable_26
Я
IsVariableInitialized_42IsVariableInitializedtraining/Adam/Variable_27*,
_class"
 loc:@training/Adam/Variable_27*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_43IsVariableInitializedtraining/Adam/Variable_28*,
_class"
 loc:@training/Adam/Variable_28*
dtype0*
_output_shapes
: 
Я
IsVariableInitialized_44IsVariableInitializedtraining/Adam/Variable_29*
_output_shapes
: *,
_class"
 loc:@training/Adam/Variable_29*
dtype0
·

initNoOp^Adam/beta_1/Assign^Adam/beta_2/Assign^Adam/decay/Assign^Adam/iterations/Assign^Adam/lr/Assign^dense_1/bias/Assign^dense_1/kernel/Assign^dense_2/bias/Assign^dense_2/kernel/Assign^dense_3/bias/Assign^dense_3/kernel/Assign^input_layer/bias/Assign^input_layer/kernel/Assign^output_layer/bias/Assign^output_layer/kernel/Assign^training/Adam/Variable/Assign ^training/Adam/Variable_1/Assign!^training/Adam/Variable_10/Assign!^training/Adam/Variable_11/Assign!^training/Adam/Variable_12/Assign!^training/Adam/Variable_13/Assign!^training/Adam/Variable_14/Assign!^training/Adam/Variable_15/Assign!^training/Adam/Variable_16/Assign!^training/Adam/Variable_17/Assign!^training/Adam/Variable_18/Assign!^training/Adam/Variable_19/Assign ^training/Adam/Variable_2/Assign!^training/Adam/Variable_20/Assign!^training/Adam/Variable_21/Assign!^training/Adam/Variable_22/Assign!^training/Adam/Variable_23/Assign!^training/Adam/Variable_24/Assign!^training/Adam/Variable_25/Assign!^training/Adam/Variable_26/Assign!^training/Adam/Variable_27/Assign!^training/Adam/Variable_28/Assign!^training/Adam/Variable_29/Assign ^training/Adam/Variable_3/Assign ^training/Adam/Variable_4/Assign ^training/Adam/Variable_5/Assign ^training/Adam/Variable_6/Assign ^training/Adam/Variable_7/Assign ^training/Adam/Variable_8/Assign ^training/Adam/Variable_9/Assign
№

init_1NoOp^Adam/beta_1/Assign^Adam/beta_2/Assign^Adam/decay/Assign^Adam/iterations/Assign^Adam/lr/Assign^dense_1/bias/Assign^dense_1/kernel/Assign^dense_2/bias/Assign^dense_2/kernel/Assign^dense_3/bias/Assign^dense_3/kernel/Assign^input_layer/bias/Assign^input_layer/kernel/Assign^output_layer/bias/Assign^output_layer/kernel/Assign^training/Adam/Variable/Assign ^training/Adam/Variable_1/Assign!^training/Adam/Variable_10/Assign!^training/Adam/Variable_11/Assign!^training/Adam/Variable_12/Assign!^training/Adam/Variable_13/Assign!^training/Adam/Variable_14/Assign!^training/Adam/Variable_15/Assign!^training/Adam/Variable_16/Assign!^training/Adam/Variable_17/Assign!^training/Adam/Variable_18/Assign!^training/Adam/Variable_19/Assign ^training/Adam/Variable_2/Assign!^training/Adam/Variable_20/Assign!^training/Adam/Variable_21/Assign!^training/Adam/Variable_22/Assign!^training/Adam/Variable_23/Assign!^training/Adam/Variable_24/Assign!^training/Adam/Variable_25/Assign!^training/Adam/Variable_26/Assign!^training/Adam/Variable_27/Assign!^training/Adam/Variable_28/Assign!^training/Adam/Variable_29/Assign ^training/Adam/Variable_3/Assign ^training/Adam/Variable_4/Assign ^training/Adam/Variable_5/Assign ^training/Adam/Variable_6/Assign ^training/Adam/Variable_7/Assign ^training/Adam/Variable_8/Assign ^training/Adam/Variable_9/Assign
a
Adam_1/iterations/initial_valueConst*
value	B	 R *
dtype0	*
_output_shapes
: 
u
Adam_1/iterations
VariableV2*
_output_shapes
: *
	container *
shape: *
shared_name *
dtype0	
╞
Adam_1/iterations/AssignAssignAdam_1/iterationsAdam_1/iterations/initial_value*
validate_shape(*
_output_shapes
: *
use_locking(*
T0	*$
_class
loc:@Adam_1/iterations
|
Adam_1/iterations/readIdentityAdam_1/iterations*
_output_shapes
: *
T0	*$
_class
loc:@Adam_1/iterations
\
Adam_1/lr/initial_valueConst*
valueB
 *╖╤8*
dtype0*
_output_shapes
: 
m
	Adam_1/lr
VariableV2*
shape: *
shared_name *
dtype0*
_output_shapes
: *
	container 
ж
Adam_1/lr/AssignAssign	Adam_1/lrAdam_1/lr/initial_value*
T0*
_class
loc:@Adam_1/lr*
validate_shape(*
_output_shapes
: *
use_locking(
d
Adam_1/lr/readIdentity	Adam_1/lr*
T0*
_class
loc:@Adam_1/lr*
_output_shapes
: 
`
Adam_1/beta_1/initial_valueConst*
valueB
 *fff?*
dtype0*
_output_shapes
: 
q
Adam_1/beta_1
VariableV2*
_output_shapes
: *
	container *
shape: *
shared_name *
dtype0
╢
Adam_1/beta_1/AssignAssignAdam_1/beta_1Adam_1/beta_1/initial_value* 
_class
loc:@Adam_1/beta_1*
validate_shape(*
_output_shapes
: *
use_locking(*
T0
p
Adam_1/beta_1/readIdentityAdam_1/beta_1* 
_class
loc:@Adam_1/beta_1*
_output_shapes
: *
T0
`
Adam_1/beta_2/initial_valueConst*
dtype0*
_output_shapes
: *
valueB
 *w╛?
q
Adam_1/beta_2
VariableV2*
_output_shapes
: *
	container *
shape: *
shared_name *
dtype0
╢
Adam_1/beta_2/AssignAssignAdam_1/beta_2Adam_1/beta_2/initial_value*
use_locking(*
T0* 
_class
loc:@Adam_1/beta_2*
validate_shape(*
_output_shapes
: 
p
Adam_1/beta_2/readIdentityAdam_1/beta_2*
T0* 
_class
loc:@Adam_1/beta_2*
_output_shapes
: 
_
Adam_1/decay/initial_valueConst*
dtype0*
_output_shapes
: *
valueB
 *    
p
Adam_1/decay
VariableV2*
shape: *
shared_name *
dtype0*
_output_shapes
: *
	container 
▓
Adam_1/decay/AssignAssignAdam_1/decayAdam_1/decay/initial_value*
use_locking(*
T0*
_class
loc:@Adam_1/decay*
validate_shape(*
_output_shapes
: 
m
Adam_1/decay/readIdentityAdam_1/decay*
T0*
_class
loc:@Adam_1/decay*
_output_shapes
: 
^
PlaceholderPlaceholder*
dtype0*
_output_shapes
:	zА*
shape:	zА
л
AssignAssigninput_layer/kernelPlaceholder*
use_locking( *
T0*%
_class
loc:@input_layer/kernel*
validate_shape(*
_output_shapes
:	zА
X
Placeholder_1Placeholder*
_output_shapes	
:А*
shape:А*
dtype0
з
Assign_1Assigninput_layer/biasPlaceholder_1*
_output_shapes	
:А*
use_locking( *
T0*#
_class
loc:@input_layer/bias*
validate_shape(
b
Placeholder_2Placeholder*
dtype0* 
_output_shapes
:
АА*
shape:
АА
и
Assign_2Assigndense_1/kernelPlaceholder_2*
use_locking( *
T0*!
_class
loc:@dense_1/kernel*
validate_shape(* 
_output_shapes
:
АА
X
Placeholder_3Placeholder*
dtype0*
_output_shapes	
:А*
shape:А
Я
Assign_3Assigndense_1/biasPlaceholder_3*
T0*
_class
loc:@dense_1/bias*
validate_shape(*
_output_shapes	
:А*
use_locking( 
b
Placeholder_4Placeholder*
shape:
АА*
dtype0* 
_output_shapes
:
АА
и
Assign_4Assigndense_2/kernelPlaceholder_4*!
_class
loc:@dense_2/kernel*
validate_shape(* 
_output_shapes
:
АА*
use_locking( *
T0
X
Placeholder_5Placeholder*
dtype0*
_output_shapes	
:А*
shape:А
Я
Assign_5Assigndense_2/biasPlaceholder_5*
validate_shape(*
_output_shapes	
:А*
use_locking( *
T0*
_class
loc:@dense_2/bias
b
Placeholder_6Placeholder*
shape:
АА*
dtype0* 
_output_shapes
:
АА
и
Assign_6Assigndense_3/kernelPlaceholder_6* 
_output_shapes
:
АА*
use_locking( *
T0*!
_class
loc:@dense_3/kernel*
validate_shape(
X
Placeholder_7Placeholder*
dtype0*
_output_shapes	
:А*
shape:А
Я
Assign_7Assigndense_3/biasPlaceholder_7*
use_locking( *
T0*
_class
loc:@dense_3/bias*
validate_shape(*
_output_shapes	
:А
`
Placeholder_8Placeholder*
dtype0*
_output_shapes
:	А(*
shape:	А(
▒
Assign_8Assignoutput_layer/kernelPlaceholder_8*
T0*&
_class
loc:@output_layer/kernel*
validate_shape(*
_output_shapes
:	А(*
use_locking( 
V
Placeholder_9Placeholder*
dtype0*
_output_shapes
:(*
shape:(
и
Assign_9Assignoutput_layer/biasPlaceholder_9*$
_class
loc:@output_layer/bias*
validate_shape(*
_output_shapes
:(*
use_locking( *
T0
П
IsVariableInitialized_45IsVariableInitializedAdam_1/iterations*$
_class
loc:@Adam_1/iterations*
dtype0	*
_output_shapes
: 

IsVariableInitialized_46IsVariableInitialized	Adam_1/lr*
_class
loc:@Adam_1/lr*
dtype0*
_output_shapes
: 
З
IsVariableInitialized_47IsVariableInitializedAdam_1/beta_1* 
_class
loc:@Adam_1/beta_1*
dtype0*
_output_shapes
: 
З
IsVariableInitialized_48IsVariableInitializedAdam_1/beta_2* 
_class
loc:@Adam_1/beta_2*
dtype0*
_output_shapes
: 
Е
IsVariableInitialized_49IsVariableInitializedAdam_1/decay*
_class
loc:@Adam_1/decay*
dtype0*
_output_shapes
: 
А
init_2NoOp^Adam_1/beta_1/Assign^Adam_1/beta_2/Assign^Adam_1/decay/Assign^Adam_1/iterations/Assign^Adam_1/lr/Assign
д
output_layer_target_1Placeholder*
dtype0*=
_output_shapes+
):'                           *2
shape):'                           
x
output_layer_sample_weights_1Placeholder*
shape:         *
dtype0*#
_output_shapes
:         
Л
loss_1/output_layer_loss/subSuboutput_layer/addoutput_layer_target_1*
T0*4
_output_shapes"
 :                  (
А
loss_1/output_layer_loss/AbsAbsloss_1/output_layer_loss/sub*
T0*4
_output_shapes"
 :                  (
z
/loss_1/output_layer_loss/Mean/reduction_indicesConst*
valueB :
         *
dtype0*
_output_shapes
: 
╠
loss_1/output_layer_loss/MeanMeanloss_1/output_layer_loss/Abs/loss_1/output_layer_loss/Mean/reduction_indices*
T0*0
_output_shapes
:                  *
	keep_dims( *

Tidx0
{
1loss_1/output_layer_loss/Mean_1/reduction_indicesConst*
valueB:*
dtype0*
_output_shapes
:
─
loss_1/output_layer_loss/Mean_1Meanloss_1/output_layer_loss/Mean1loss_1/output_layer_loss/Mean_1/reduction_indices*#
_output_shapes
:         *
	keep_dims( *

Tidx0*
T0
С
loss_1/output_layer_loss/mulMulloss_1/output_layer_loss/Mean_1output_layer_sample_weights_1*
T0*#
_output_shapes
:         
h
#loss_1/output_layer_loss/NotEqual/yConst*
dtype0*
_output_shapes
: *
valueB
 *    
Я
!loss_1/output_layer_loss/NotEqualNotEqualoutput_layer_sample_weights_1#loss_1/output_layer_loss/NotEqual/y*
T0*#
_output_shapes
:         
Х
loss_1/output_layer_loss/CastCast!loss_1/output_layer_loss/NotEqual*
Truncate( *#
_output_shapes
:         *

DstT0*

SrcT0

h
loss_1/output_layer_loss/ConstConst*
valueB: *
dtype0*
_output_shapes
:
д
loss_1/output_layer_loss/Mean_2Meanloss_1/output_layer_loss/Castloss_1/output_layer_loss/Const*
T0*
_output_shapes
: *
	keep_dims( *

Tidx0
Ш
 loss_1/output_layer_loss/truedivRealDivloss_1/output_layer_loss/mulloss_1/output_layer_loss/Mean_2*
T0*#
_output_shapes
:         
j
 loss_1/output_layer_loss/Const_1Const*
valueB: *
dtype0*
_output_shapes
:
й
loss_1/output_layer_loss/Mean_3Mean loss_1/output_layer_loss/truediv loss_1/output_layer_loss/Const_1*
_output_shapes
: *
	keep_dims( *

Tidx0*
T0
Q
loss_1/mul/xConst*
_output_shapes
: *
valueB
 *  А?*
dtype0
a

loss_1/mulMulloss_1/mul/xloss_1/output_layer_loss/Mean_3*
T0*
_output_shapes
: 
Р
!metrics_1/mean_absolute_error/subSuboutput_layer/addoutput_layer_target_1*
T0*4
_output_shapes"
 :                  (
К
!metrics_1/mean_absolute_error/AbsAbs!metrics_1/mean_absolute_error/sub*
T0*4
_output_shapes"
 :                  (

4metrics_1/mean_absolute_error/Mean/reduction_indicesConst*
valueB :
         *
dtype0*
_output_shapes
: 
█
"metrics_1/mean_absolute_error/MeanMean!metrics_1/mean_absolute_error/Abs4metrics_1/mean_absolute_error/Mean/reduction_indices*0
_output_shapes
:                  *
	keep_dims( *

Tidx0*
T0
t
#metrics_1/mean_absolute_error/ConstConst*
valueB"       *
dtype0*
_output_shapes
:
│
$metrics_1/mean_absolute_error/Mean_1Mean"metrics_1/mean_absolute_error/Mean#metrics_1/mean_absolute_error/Const*
	keep_dims( *

Tidx0*
T0*
_output_shapes
: 
'
group_deps_1NoOp^output_layer/add
Y
save/filename/inputConst*
dtype0*
_output_shapes
: *
valueB Bmodel
n
save/filenamePlaceholderWithDefaultsave/filename/input*
dtype0*
_output_shapes
: *
shape: 
e

save/ConstPlaceholderWithDefaultsave/filename*
_output_shapes
: *
shape: *
dtype0
Д
save/StringJoin/inputs_1Const*<
value3B1 B+_temp_651b849f8f794b9987e7a355db432475/part*
dtype0*
_output_shapes
: 
u
save/StringJoin
StringJoin
save/Constsave/StringJoin/inputs_1*
	separator *
N*
_output_shapes
: 
Q
save/num_shardsConst*
value	B :*
dtype0*
_output_shapes
: 
\
save/ShardedFilename/shardConst*
value	B : *
dtype0*
_output_shapes
: 
}
save/ShardedFilenameShardedFilenamesave/StringJoinsave/ShardedFilename/shardsave/num_shards*
_output_shapes
: 
▓	
save/SaveV2/tensor_namesConst*х
value█B╪2BAdam/beta_1BAdam/beta_2B
Adam/decayBAdam/iterationsBAdam/lrBAdam_1/beta_1BAdam_1/beta_2BAdam_1/decayBAdam_1/iterationsB	Adam_1/lrBdense_1/biasBdense_1/kernelBdense_2/biasBdense_2/kernelBdense_3/biasBdense_3/kernelBinput_layer/biasBinput_layer/kernelBoutput_layer/biasBoutput_layer/kernelBtraining/Adam/VariableBtraining/Adam/Variable_1Btraining/Adam/Variable_10Btraining/Adam/Variable_11Btraining/Adam/Variable_12Btraining/Adam/Variable_13Btraining/Adam/Variable_14Btraining/Adam/Variable_15Btraining/Adam/Variable_16Btraining/Adam/Variable_17Btraining/Adam/Variable_18Btraining/Adam/Variable_19Btraining/Adam/Variable_2Btraining/Adam/Variable_20Btraining/Adam/Variable_21Btraining/Adam/Variable_22Btraining/Adam/Variable_23Btraining/Adam/Variable_24Btraining/Adam/Variable_25Btraining/Adam/Variable_26Btraining/Adam/Variable_27Btraining/Adam/Variable_28Btraining/Adam/Variable_29Btraining/Adam/Variable_3Btraining/Adam/Variable_4Btraining/Adam/Variable_5Btraining/Adam/Variable_6Btraining/Adam/Variable_7Btraining/Adam/Variable_8Btraining/Adam/Variable_9*
dtype0*
_output_shapes
:2
╟
save/SaveV2/shape_and_slicesConst*w
valuenBl2B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B *
dtype0*
_output_shapes
:2
ї	
save/SaveV2SaveV2save/ShardedFilenamesave/SaveV2/tensor_namessave/SaveV2/shape_and_slicesAdam/beta_1Adam/beta_2
Adam/decayAdam/iterationsAdam/lrAdam_1/beta_1Adam_1/beta_2Adam_1/decayAdam_1/iterations	Adam_1/lrdense_1/biasdense_1/kerneldense_2/biasdense_2/kerneldense_3/biasdense_3/kernelinput_layer/biasinput_layer/kerneloutput_layer/biasoutput_layer/kerneltraining/Adam/Variabletraining/Adam/Variable_1training/Adam/Variable_10training/Adam/Variable_11training/Adam/Variable_12training/Adam/Variable_13training/Adam/Variable_14training/Adam/Variable_15training/Adam/Variable_16training/Adam/Variable_17training/Adam/Variable_18training/Adam/Variable_19training/Adam/Variable_2training/Adam/Variable_20training/Adam/Variable_21training/Adam/Variable_22training/Adam/Variable_23training/Adam/Variable_24training/Adam/Variable_25training/Adam/Variable_26training/Adam/Variable_27training/Adam/Variable_28training/Adam/Variable_29training/Adam/Variable_3training/Adam/Variable_4training/Adam/Variable_5training/Adam/Variable_6training/Adam/Variable_7training/Adam/Variable_8training/Adam/Variable_9*@
dtypes6
422		
С
save/control_dependencyIdentitysave/ShardedFilename^save/SaveV2*
T0*'
_class
loc:@save/ShardedFilename*
_output_shapes
: 
Э
+save/MergeV2Checkpoints/checkpoint_prefixesPacksave/ShardedFilename^save/control_dependency*
T0*

axis *
N*
_output_shapes
:
}
save/MergeV2CheckpointsMergeV2Checkpoints+save/MergeV2Checkpoints/checkpoint_prefixes
save/Const*
delete_old_dirs(
z
save/IdentityIdentity
save/Const^save/MergeV2Checkpoints^save/control_dependency*
T0*
_output_shapes
: 
╡	
save/RestoreV2/tensor_namesConst*х
value█B╪2BAdam/beta_1BAdam/beta_2B
Adam/decayBAdam/iterationsBAdam/lrBAdam_1/beta_1BAdam_1/beta_2BAdam_1/decayBAdam_1/iterationsB	Adam_1/lrBdense_1/biasBdense_1/kernelBdense_2/biasBdense_2/kernelBdense_3/biasBdense_3/kernelBinput_layer/biasBinput_layer/kernelBoutput_layer/biasBoutput_layer/kernelBtraining/Adam/VariableBtraining/Adam/Variable_1Btraining/Adam/Variable_10Btraining/Adam/Variable_11Btraining/Adam/Variable_12Btraining/Adam/Variable_13Btraining/Adam/Variable_14Btraining/Adam/Variable_15Btraining/Adam/Variable_16Btraining/Adam/Variable_17Btraining/Adam/Variable_18Btraining/Adam/Variable_19Btraining/Adam/Variable_2Btraining/Adam/Variable_20Btraining/Adam/Variable_21Btraining/Adam/Variable_22Btraining/Adam/Variable_23Btraining/Adam/Variable_24Btraining/Adam/Variable_25Btraining/Adam/Variable_26Btraining/Adam/Variable_27Btraining/Adam/Variable_28Btraining/Adam/Variable_29Btraining/Adam/Variable_3Btraining/Adam/Variable_4Btraining/Adam/Variable_5Btraining/Adam/Variable_6Btraining/Adam/Variable_7Btraining/Adam/Variable_8Btraining/Adam/Variable_9*
dtype0*
_output_shapes
:2
╩
save/RestoreV2/shape_and_slicesConst*
_output_shapes
:2*w
valuenBl2B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B B *
dtype0
И
save/RestoreV2	RestoreV2
save/Constsave/RestoreV2/tensor_namessave/RestoreV2/shape_and_slices*▐
_output_shapes╦
╚::::::::::::::::::::::::::::::::::::::::::::::::::*@
dtypes6
422		
Ь
save/AssignAssignAdam/beta_1save/RestoreV2*
_class
loc:@Adam/beta_1*
validate_shape(*
_output_shapes
: *
use_locking(*
T0
а
save/Assign_1AssignAdam/beta_2save/RestoreV2:1*
_output_shapes
: *
use_locking(*
T0*
_class
loc:@Adam/beta_2*
validate_shape(
Ю
save/Assign_2Assign
Adam/decaysave/RestoreV2:2*
_class
loc:@Adam/decay*
validate_shape(*
_output_shapes
: *
use_locking(*
T0
и
save/Assign_3AssignAdam/iterationssave/RestoreV2:3*
use_locking(*
T0	*"
_class
loc:@Adam/iterations*
validate_shape(*
_output_shapes
: 
Ш
save/Assign_4AssignAdam/lrsave/RestoreV2:4*
_class
loc:@Adam/lr*
validate_shape(*
_output_shapes
: *
use_locking(*
T0
д
save/Assign_5AssignAdam_1/beta_1save/RestoreV2:5*
use_locking(*
T0* 
_class
loc:@Adam_1/beta_1*
validate_shape(*
_output_shapes
: 
д
save/Assign_6AssignAdam_1/beta_2save/RestoreV2:6*
use_locking(*
T0* 
_class
loc:@Adam_1/beta_2*
validate_shape(*
_output_shapes
: 
в
save/Assign_7AssignAdam_1/decaysave/RestoreV2:7*
_output_shapes
: *
use_locking(*
T0*
_class
loc:@Adam_1/decay*
validate_shape(
м
save/Assign_8AssignAdam_1/iterationssave/RestoreV2:8*
T0	*$
_class
loc:@Adam_1/iterations*
validate_shape(*
_output_shapes
: *
use_locking(
Ь
save/Assign_9Assign	Adam_1/lrsave/RestoreV2:9*
use_locking(*
T0*
_class
loc:@Adam_1/lr*
validate_shape(*
_output_shapes
: 
й
save/Assign_10Assigndense_1/biassave/RestoreV2:10*
validate_shape(*
_output_shapes	
:А*
use_locking(*
T0*
_class
loc:@dense_1/bias
▓
save/Assign_11Assigndense_1/kernelsave/RestoreV2:11* 
_output_shapes
:
АА*
use_locking(*
T0*!
_class
loc:@dense_1/kernel*
validate_shape(
й
save/Assign_12Assigndense_2/biassave/RestoreV2:12*
use_locking(*
T0*
_class
loc:@dense_2/bias*
validate_shape(*
_output_shapes	
:А
▓
save/Assign_13Assigndense_2/kernelsave/RestoreV2:13*
use_locking(*
T0*!
_class
loc:@dense_2/kernel*
validate_shape(* 
_output_shapes
:
АА
й
save/Assign_14Assigndense_3/biassave/RestoreV2:14*
use_locking(*
T0*
_class
loc:@dense_3/bias*
validate_shape(*
_output_shapes	
:А
▓
save/Assign_15Assigndense_3/kernelsave/RestoreV2:15*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0*!
_class
loc:@dense_3/kernel
▒
save/Assign_16Assigninput_layer/biassave/RestoreV2:16*
T0*#
_class
loc:@input_layer/bias*
validate_shape(*
_output_shapes	
:А*
use_locking(
╣
save/Assign_17Assigninput_layer/kernelsave/RestoreV2:17*
use_locking(*
T0*%
_class
loc:@input_layer/kernel*
validate_shape(*
_output_shapes
:	zА
▓
save/Assign_18Assignoutput_layer/biassave/RestoreV2:18*
_output_shapes
:(*
use_locking(*
T0*$
_class
loc:@output_layer/bias*
validate_shape(
╗
save/Assign_19Assignoutput_layer/kernelsave/RestoreV2:19*
use_locking(*
T0*&
_class
loc:@output_layer/kernel*
validate_shape(*
_output_shapes
:	А(
┴
save/Assign_20Assigntraining/Adam/Variablesave/RestoreV2:20*)
_class
loc:@training/Adam/Variable*
validate_shape(*
_output_shapes
:	zА*
use_locking(*
T0
┴
save/Assign_21Assigntraining/Adam/Variable_1save/RestoreV2:21*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_1*
validate_shape(*
_output_shapes	
:А
╟
save/Assign_22Assigntraining/Adam/Variable_10save/RestoreV2:22*
_output_shapes
:	zА*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_10*
validate_shape(
├
save/Assign_23Assigntraining/Adam/Variable_11save/RestoreV2:23*
validate_shape(*
_output_shapes	
:А*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_11
╚
save/Assign_24Assigntraining/Adam/Variable_12save/RestoreV2:24*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_12
├
save/Assign_25Assigntraining/Adam/Variable_13save/RestoreV2:25*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_13*
validate_shape(*
_output_shapes	
:А
╚
save/Assign_26Assigntraining/Adam/Variable_14save/RestoreV2:26*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_14*
validate_shape(* 
_output_shapes
:
АА
├
save/Assign_27Assigntraining/Adam/Variable_15save/RestoreV2:27*
validate_shape(*
_output_shapes	
:А*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_15
╚
save/Assign_28Assigntraining/Adam/Variable_16save/RestoreV2:28*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_16
├
save/Assign_29Assigntraining/Adam/Variable_17save/RestoreV2:29*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_17*
validate_shape(*
_output_shapes	
:А
╟
save/Assign_30Assigntraining/Adam/Variable_18save/RestoreV2:30*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_18*
validate_shape(*
_output_shapes
:	А(
┬
save/Assign_31Assigntraining/Adam/Variable_19save/RestoreV2:31*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_19*
validate_shape(*
_output_shapes
:(
╞
save/Assign_32Assigntraining/Adam/Variable_2save/RestoreV2:32*+
_class!
loc:@training/Adam/Variable_2*
validate_shape(* 
_output_shapes
:
АА*
use_locking(*
T0
┬
save/Assign_33Assigntraining/Adam/Variable_20save/RestoreV2:33*,
_class"
 loc:@training/Adam/Variable_20*
validate_shape(*
_output_shapes
:*
use_locking(*
T0
┬
save/Assign_34Assigntraining/Adam/Variable_21save/RestoreV2:34*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_21*
validate_shape(*
_output_shapes
:
┬
save/Assign_35Assigntraining/Adam/Variable_22save/RestoreV2:35*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_22*
validate_shape(*
_output_shapes
:
┬
save/Assign_36Assigntraining/Adam/Variable_23save/RestoreV2:36*,
_class"
 loc:@training/Adam/Variable_23*
validate_shape(*
_output_shapes
:*
use_locking(*
T0
┬
save/Assign_37Assigntraining/Adam/Variable_24save/RestoreV2:37*,
_class"
 loc:@training/Adam/Variable_24*
validate_shape(*
_output_shapes
:*
use_locking(*
T0
┬
save/Assign_38Assigntraining/Adam/Variable_25save/RestoreV2:38*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_25*
validate_shape(*
_output_shapes
:
┬
save/Assign_39Assigntraining/Adam/Variable_26save/RestoreV2:39*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_26*
validate_shape(*
_output_shapes
:
┬
save/Assign_40Assigntraining/Adam/Variable_27save/RestoreV2:40*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_27*
validate_shape(*
_output_shapes
:
┬
save/Assign_41Assigntraining/Adam/Variable_28save/RestoreV2:41*
validate_shape(*
_output_shapes
:*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_28
┬
save/Assign_42Assigntraining/Adam/Variable_29save/RestoreV2:42*
use_locking(*
T0*,
_class"
 loc:@training/Adam/Variable_29*
validate_shape(*
_output_shapes
:
┴
save/Assign_43Assigntraining/Adam/Variable_3save/RestoreV2:43*
T0*+
_class!
loc:@training/Adam/Variable_3*
validate_shape(*
_output_shapes	
:А*
use_locking(
╞
save/Assign_44Assigntraining/Adam/Variable_4save/RestoreV2:44*
T0*+
_class!
loc:@training/Adam/Variable_4*
validate_shape(* 
_output_shapes
:
АА*
use_locking(
┴
save/Assign_45Assigntraining/Adam/Variable_5save/RestoreV2:45*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_5*
validate_shape(*
_output_shapes	
:А
╞
save/Assign_46Assigntraining/Adam/Variable_6save/RestoreV2:46*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_6*
validate_shape(* 
_output_shapes
:
АА
┴
save/Assign_47Assigntraining/Adam/Variable_7save/RestoreV2:47*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_7*
validate_shape(*
_output_shapes	
:А
┼
save/Assign_48Assigntraining/Adam/Variable_8save/RestoreV2:48*
validate_shape(*
_output_shapes
:	А(*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_8
└
save/Assign_49Assigntraining/Adam/Variable_9save/RestoreV2:49*
_output_shapes
:(*
use_locking(*
T0*+
_class!
loc:@training/Adam/Variable_9*
validate_shape(
р
save/restore_shardNoOp^save/Assign^save/Assign_1^save/Assign_10^save/Assign_11^save/Assign_12^save/Assign_13^save/Assign_14^save/Assign_15^save/Assign_16^save/Assign_17^save/Assign_18^save/Assign_19^save/Assign_2^save/Assign_20^save/Assign_21^save/Assign_22^save/Assign_23^save/Assign_24^save/Assign_25^save/Assign_26^save/Assign_27^save/Assign_28^save/Assign_29^save/Assign_3^save/Assign_30^save/Assign_31^save/Assign_32^save/Assign_33^save/Assign_34^save/Assign_35^save/Assign_36^save/Assign_37^save/Assign_38^save/Assign_39^save/Assign_4^save/Assign_40^save/Assign_41^save/Assign_42^save/Assign_43^save/Assign_44^save/Assign_45^save/Assign_46^save/Assign_47^save/Assign_48^save/Assign_49^save/Assign_5^save/Assign_6^save/Assign_7^save/Assign_8^save/Assign_9
-
save/restore_allNoOp^save/restore_shard "<
save/Const:0save/Identity:0save/restore_all (5 @F8"Б,
trainable_variablesщ+ц+
k
input_layer/kernel:0input_layer/kernel/Assigninput_layer/kernel/read:02input_layer/random_normal:08
]
input_layer/bias:0input_layer/bias/Assigninput_layer/bias/read:02input_layer/Const:08
[
dense_1/kernel:0dense_1/kernel/Assigndense_1/kernel/read:02dense_1/random_normal:08
M
dense_1/bias:0dense_1/bias/Assigndense_1/bias/read:02dense_1/Const:08
[
dense_2/kernel:0dense_2/kernel/Assigndense_2/kernel/read:02dense_2/random_normal:08
M
dense_2/bias:0dense_2/bias/Assigndense_2/bias/read:02dense_2/Const:08
[
dense_3/kernel:0dense_3/kernel/Assigndense_3/kernel/read:02dense_3/random_normal:08
M
dense_3/bias:0dense_3/bias/Assigndense_3/bias/read:02dense_3/Const:08
o
output_layer/kernel:0output_layer/kernel/Assignoutput_layer/kernel/read:02output_layer/random_normal:08
a
output_layer/bias:0output_layer/bias/Assignoutput_layer/bias/read:02output_layer/Const:08
f
Adam/iterations:0Adam/iterations/AssignAdam/iterations/read:02Adam/iterations/initial_value:08
F
	Adam/lr:0Adam/lr/AssignAdam/lr/read:02Adam/lr/initial_value:08
V
Adam/beta_1:0Adam/beta_1/AssignAdam/beta_1/read:02Adam/beta_1/initial_value:08
V
Adam/beta_2:0Adam/beta_2/AssignAdam/beta_2/read:02Adam/beta_2/initial_value:08
R
Adam/decay:0Adam/decay/AssignAdam/decay/read:02Adam/decay/initial_value:08
q
training/Adam/Variable:0training/Adam/Variable/Assigntraining/Adam/Variable/read:02training/Adam/zeros:08
y
training/Adam/Variable_1:0training/Adam/Variable_1/Assigntraining/Adam/Variable_1/read:02training/Adam/zeros_1:08
y
training/Adam/Variable_2:0training/Adam/Variable_2/Assigntraining/Adam/Variable_2/read:02training/Adam/zeros_2:08
y
training/Adam/Variable_3:0training/Adam/Variable_3/Assigntraining/Adam/Variable_3/read:02training/Adam/zeros_3:08
y
training/Adam/Variable_4:0training/Adam/Variable_4/Assigntraining/Adam/Variable_4/read:02training/Adam/zeros_4:08
y
training/Adam/Variable_5:0training/Adam/Variable_5/Assigntraining/Adam/Variable_5/read:02training/Adam/zeros_5:08
y
training/Adam/Variable_6:0training/Adam/Variable_6/Assigntraining/Adam/Variable_6/read:02training/Adam/zeros_6:08
y
training/Adam/Variable_7:0training/Adam/Variable_7/Assigntraining/Adam/Variable_7/read:02training/Adam/zeros_7:08
y
training/Adam/Variable_8:0training/Adam/Variable_8/Assigntraining/Adam/Variable_8/read:02training/Adam/zeros_8:08
y
training/Adam/Variable_9:0training/Adam/Variable_9/Assigntraining/Adam/Variable_9/read:02training/Adam/zeros_9:08
}
training/Adam/Variable_10:0 training/Adam/Variable_10/Assign training/Adam/Variable_10/read:02training/Adam/zeros_10:08
}
training/Adam/Variable_11:0 training/Adam/Variable_11/Assign training/Adam/Variable_11/read:02training/Adam/zeros_11:08
}
training/Adam/Variable_12:0 training/Adam/Variable_12/Assign training/Adam/Variable_12/read:02training/Adam/zeros_12:08
}
training/Adam/Variable_13:0 training/Adam/Variable_13/Assign training/Adam/Variable_13/read:02training/Adam/zeros_13:08
}
training/Adam/Variable_14:0 training/Adam/Variable_14/Assign training/Adam/Variable_14/read:02training/Adam/zeros_14:08
}
training/Adam/Variable_15:0 training/Adam/Variable_15/Assign training/Adam/Variable_15/read:02training/Adam/zeros_15:08
}
training/Adam/Variable_16:0 training/Adam/Variable_16/Assign training/Adam/Variable_16/read:02training/Adam/zeros_16:08
}
training/Adam/Variable_17:0 training/Adam/Variable_17/Assign training/Adam/Variable_17/read:02training/Adam/zeros_17:08
}
training/Adam/Variable_18:0 training/Adam/Variable_18/Assign training/Adam/Variable_18/read:02training/Adam/zeros_18:08
}
training/Adam/Variable_19:0 training/Adam/Variable_19/Assign training/Adam/Variable_19/read:02training/Adam/zeros_19:08
}
training/Adam/Variable_20:0 training/Adam/Variable_20/Assign training/Adam/Variable_20/read:02training/Adam/zeros_20:08
}
training/Adam/Variable_21:0 training/Adam/Variable_21/Assign training/Adam/Variable_21/read:02training/Adam/zeros_21:08
}
training/Adam/Variable_22:0 training/Adam/Variable_22/Assign training/Adam/Variable_22/read:02training/Adam/zeros_22:08
}
training/Adam/Variable_23:0 training/Adam/Variable_23/Assign training/Adam/Variable_23/read:02training/Adam/zeros_23:08
}
training/Adam/Variable_24:0 training/Adam/Variable_24/Assign training/Adam/Variable_24/read:02training/Adam/zeros_24:08
}
training/Adam/Variable_25:0 training/Adam/Variable_25/Assign training/Adam/Variable_25/read:02training/Adam/zeros_25:08
}
training/Adam/Variable_26:0 training/Adam/Variable_26/Assign training/Adam/Variable_26/read:02training/Adam/zeros_26:08
}
training/Adam/Variable_27:0 training/Adam/Variable_27/Assign training/Adam/Variable_27/read:02training/Adam/zeros_27:08
}
training/Adam/Variable_28:0 training/Adam/Variable_28/Assign training/Adam/Variable_28/read:02training/Adam/zeros_28:08
}
training/Adam/Variable_29:0 training/Adam/Variable_29/Assign training/Adam/Variable_29/read:02training/Adam/zeros_29:08
n
Adam_1/iterations:0Adam_1/iterations/AssignAdam_1/iterations/read:02!Adam_1/iterations/initial_value:08
N
Adam_1/lr:0Adam_1/lr/AssignAdam_1/lr/read:02Adam_1/lr/initial_value:08
^
Adam_1/beta_1:0Adam_1/beta_1/AssignAdam_1/beta_1/read:02Adam_1/beta_1/initial_value:08
^
Adam_1/beta_2:0Adam_1/beta_2/AssignAdam_1/beta_2/read:02Adam_1/beta_2/initial_value:08
Z
Adam_1/decay:0Adam_1/decay/AssignAdam_1/decay/read:02Adam_1/decay/initial_value:08"ў+
	variablesщ+ц+
k
input_layer/kernel:0input_layer/kernel/Assigninput_layer/kernel/read:02input_layer/random_normal:08
]
input_layer/bias:0input_layer/bias/Assigninput_layer/bias/read:02input_layer/Const:08
[
dense_1/kernel:0dense_1/kernel/Assigndense_1/kernel/read:02dense_1/random_normal:08
M
dense_1/bias:0dense_1/bias/Assigndense_1/bias/read:02dense_1/Const:08
[
dense_2/kernel:0dense_2/kernel/Assigndense_2/kernel/read:02dense_2/random_normal:08
M
dense_2/bias:0dense_2/bias/Assigndense_2/bias/read:02dense_2/Const:08
[
dense_3/kernel:0dense_3/kernel/Assigndense_3/kernel/read:02dense_3/random_normal:08
M
dense_3/bias:0dense_3/bias/Assigndense_3/bias/read:02dense_3/Const:08
o
output_layer/kernel:0output_layer/kernel/Assignoutput_layer/kernel/read:02output_layer/random_normal:08
a
output_layer/bias:0output_layer/bias/Assignoutput_layer/bias/read:02output_layer/Const:08
f
Adam/iterations:0Adam/iterations/AssignAdam/iterations/read:02Adam/iterations/initial_value:08
F
	Adam/lr:0Adam/lr/AssignAdam/lr/read:02Adam/lr/initial_value:08
V
Adam/beta_1:0Adam/beta_1/AssignAdam/beta_1/read:02Adam/beta_1/initial_value:08
V
Adam/beta_2:0Adam/beta_2/AssignAdam/beta_2/read:02Adam/beta_2/initial_value:08
R
Adam/decay:0Adam/decay/AssignAdam/decay/read:02Adam/decay/initial_value:08
q
training/Adam/Variable:0training/Adam/Variable/Assigntraining/Adam/Variable/read:02training/Adam/zeros:08
y
training/Adam/Variable_1:0training/Adam/Variable_1/Assigntraining/Adam/Variable_1/read:02training/Adam/zeros_1:08
y
training/Adam/Variable_2:0training/Adam/Variable_2/Assigntraining/Adam/Variable_2/read:02training/Adam/zeros_2:08
y
training/Adam/Variable_3:0training/Adam/Variable_3/Assigntraining/Adam/Variable_3/read:02training/Adam/zeros_3:08
y
training/Adam/Variable_4:0training/Adam/Variable_4/Assigntraining/Adam/Variable_4/read:02training/Adam/zeros_4:08
y
training/Adam/Variable_5:0training/Adam/Variable_5/Assigntraining/Adam/Variable_5/read:02training/Adam/zeros_5:08
y
training/Adam/Variable_6:0training/Adam/Variable_6/Assigntraining/Adam/Variable_6/read:02training/Adam/zeros_6:08
y
training/Adam/Variable_7:0training/Adam/Variable_7/Assigntraining/Adam/Variable_7/read:02training/Adam/zeros_7:08
y
training/Adam/Variable_8:0training/Adam/Variable_8/Assigntraining/Adam/Variable_8/read:02training/Adam/zeros_8:08
y
training/Adam/Variable_9:0training/Adam/Variable_9/Assigntraining/Adam/Variable_9/read:02training/Adam/zeros_9:08
}
training/Adam/Variable_10:0 training/Adam/Variable_10/Assign training/Adam/Variable_10/read:02training/Adam/zeros_10:08
}
training/Adam/Variable_11:0 training/Adam/Variable_11/Assign training/Adam/Variable_11/read:02training/Adam/zeros_11:08
}
training/Adam/Variable_12:0 training/Adam/Variable_12/Assign training/Adam/Variable_12/read:02training/Adam/zeros_12:08
}
training/Adam/Variable_13:0 training/Adam/Variable_13/Assign training/Adam/Variable_13/read:02training/Adam/zeros_13:08
}
training/Adam/Variable_14:0 training/Adam/Variable_14/Assign training/Adam/Variable_14/read:02training/Adam/zeros_14:08
}
training/Adam/Variable_15:0 training/Adam/Variable_15/Assign training/Adam/Variable_15/read:02training/Adam/zeros_15:08
}
training/Adam/Variable_16:0 training/Adam/Variable_16/Assign training/Adam/Variable_16/read:02training/Adam/zeros_16:08
}
training/Adam/Variable_17:0 training/Adam/Variable_17/Assign training/Adam/Variable_17/read:02training/Adam/zeros_17:08
}
training/Adam/Variable_18:0 training/Adam/Variable_18/Assign training/Adam/Variable_18/read:02training/Adam/zeros_18:08
}
training/Adam/Variable_19:0 training/Adam/Variable_19/Assign training/Adam/Variable_19/read:02training/Adam/zeros_19:08
}
training/Adam/Variable_20:0 training/Adam/Variable_20/Assign training/Adam/Variable_20/read:02training/Adam/zeros_20:08
}
training/Adam/Variable_21:0 training/Adam/Variable_21/Assign training/Adam/Variable_21/read:02training/Adam/zeros_21:08
}
training/Adam/Variable_22:0 training/Adam/Variable_22/Assign training/Adam/Variable_22/read:02training/Adam/zeros_22:08
}
training/Adam/Variable_23:0 training/Adam/Variable_23/Assign training/Adam/Variable_23/read:02training/Adam/zeros_23:08
}
training/Adam/Variable_24:0 training/Adam/Variable_24/Assign training/Adam/Variable_24/read:02training/Adam/zeros_24:08
}
training/Adam/Variable_25:0 training/Adam/Variable_25/Assign training/Adam/Variable_25/read:02training/Adam/zeros_25:08
}
training/Adam/Variable_26:0 training/Adam/Variable_26/Assign training/Adam/Variable_26/read:02training/Adam/zeros_26:08
}
training/Adam/Variable_27:0 training/Adam/Variable_27/Assign training/Adam/Variable_27/read:02training/Adam/zeros_27:08
}
training/Adam/Variable_28:0 training/Adam/Variable_28/Assign training/Adam/Variable_28/read:02training/Adam/zeros_28:08
}
training/Adam/Variable_29:0 training/Adam/Variable_29/Assign training/Adam/Variable_29/read:02training/Adam/zeros_29:08
n
Adam_1/iterations:0Adam_1/iterations/AssignAdam_1/iterations/read:02!Adam_1/iterations/initial_value:08
N
Adam_1/lr:0Adam_1/lr/AssignAdam_1/lr/read:02Adam_1/lr/initial_value:08
^
Adam_1/beta_1:0Adam_1/beta_1/AssignAdam_1/beta_1/read:02Adam_1/beta_1/initial_value:08
^
Adam_1/beta_2:0Adam_1/beta_2/AssignAdam_1/beta_2/read:02Adam_1/beta_2/initial_value:08
Z
Adam_1/decay:0Adam_1/decay/AssignAdam_1/decay/read:02Adam_1/decay/initial_value:08