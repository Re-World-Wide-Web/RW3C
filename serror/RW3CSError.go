package serror

import "strconv"

const (
	VAR_NO_EXIST   = "Var does not exist"
	FN_NO_CALLABLE = "Function is not callable"
	ERR_RT         = "Runtime"
)

func NO_OVERLOAD_2(a string, b string, op string) string {
	return a + " " + op + " " + b + " doesn't have overload"
}

func NO_CALL_FN_ARG_DIFF(a string, a_name string, a_i int, b string) string {
	return "Cannot call this function because argument #" + strconv.FormatInt(int64(a_i), 10) + " (" + a_name + ") is a " + a + " but the value type is a " + b
}

type Error struct {
	Location []Location
	Name     string `default:"System Error"`
	Message  string
}

type Location struct {
	Line   int
	Column int
	File   string
}

func (loc Location) String() string {
	return loc.File + " " + strconv.Itoa(loc.Line) + ":" + strconv.Itoa(loc.Column)
}

func (err Error) String() string {
	str := err.Name + ": " + err.Message

	for _, pos := range err.Location {
		str += "\n\t" + pos.String()
	}

	return str
}
