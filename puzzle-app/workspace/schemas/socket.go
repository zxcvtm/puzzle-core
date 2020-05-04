package schemas
import (
   "time"
)
type (
	JoinRequest struct {
	    Room      string      `json:"room"`
	}

	MoveRequest struct {
	    Opponent    string       `json:"opponent"`
	    SortArray   []int        `json:"sortArray"`
	    Attack	    string       `json:"atack"`
	    MsgDateTime time.Time    `json:"msgDateTime"`
	}
	JoinResponse struct {
	    Id          string       `json:"id"`
	    Msg         string       `json:"msg"`
	}
)