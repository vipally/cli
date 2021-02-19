package gp

//#GOGP_FILE_BEGIN
//#GOGP_IGNORE_BEGIN ///gogp_file_begin
//
/*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and compile as normal go file
//
// This is a fake go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END ///gogp_file_begin

//#GOGP_REQUIRE(github.com/os/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN ///require begin from(github.com/os/gogp/lib/fakedef)
//#GOGP_IFDEF GOGP_DO_NOT_HAS_THIS_DEFINE__
//this is a fake types for other gp file
//#GOGP_ENDIF

//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
//#GOGP_IFDEF KEY_TYPE
type GOGPKeyType int                             //
func (this GOGPKeyType) Less(o GOGPKeyType) bool { return this < o }
func (this GOGPKeyType) Show() string            { return "" } //
//#GOGP_ENDIF //KEY_TYPE

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END ///require end from(github.com/os/gogp/lib/fakedef)

////////////////////////////////////////////////////////////////////////////////

//#GOGP_IFDEF SLICE_TYPE
// IntSlice wraps []GOGPValueType to satisfy flag.Value
type GOGPGlobalNamePrefixSlice struct {
	slice      []GOGPValueType
	hasBeenSet bool
}

//#GOGP_ENDIF //SLICE_TYPE

// GOGPGlobalNamePrefixValue define a value of type GOGPValueType
type GOGPGlobalNamePrefixValue struct {
	Value       GOGPValueType   // The value from ENV of files
	Target      *GOGPValueType  // Target set the outer value pointer
	Default     GOGPValueType   // Default value
	DefaultText string          // Default value help info
	Enums       []GOGPValueType // Enumeration of valid values
	Min, Max    GOGPValueType   // [Min,Max] range of the valid values
	hasBeenSet  bool
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
