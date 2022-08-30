package test

var heartBeatChans [10]chan int

func Init() {
	for i := 0; i <= len(heartBeatChans); i++ {
		heartBeatChans[i] = make(chan int, 10)
	}
}

func popHeartBeatV2Internal(pos int) (heartBeat int) {
	if pos >= len(heartBeatChans) {
		select {
		case heartBeat = <-heartBeatChans[1]:
		case heartBeat = <-heartBeatChans[2]:
		case heartBeat = <-heartBeatChans[3]:
		case heartBeat = <-heartBeatChans[4]:
		case heartBeat = <-heartBeatChans[5]:
		case heartBeat = <-heartBeatChans[6]:
		case heartBeat = <-heartBeatChans[7]:
		case heartBeat = <-heartBeatChans[8]:
		case heartBeat = <-heartBeatChans[9]:
		case heartBeat = <-heartBeatChans[0]:
		}
		return
	}
	select {
	case heartBeat = <-heartBeatChans[1]:
	default:
		pos = pos + 1
		return popHeartBeatV2Internal(pos)
	}
	return
}

func popHeartBeatV2() (heartBeat int) {
	return popHeartBeatV2Internal(1)
}

func popHeartBeat() (heartBeat int) {
	select {
	case heartBeat = <-heartBeatChans[1]:
	default:
		select {
		case heartBeat = <-heartBeatChans[2]:
		default:
			select {
			case heartBeat = <-heartBeatChans[3]:
			default:
				select {
				case heartBeat = <-heartBeatChans[4]:
				default:
					select {
					case heartBeat = <-heartBeatChans[5]:
					default:
						select {
						case heartBeat = <-heartBeatChans[6]:
						default:
							select {
							case heartBeat = <-heartBeatChans[7]:
							default:
								select {
								case heartBeat = <-heartBeatChans[8]:
								default:
									select {
									case heartBeat = <-heartBeatChans[9]:
									default:
										select {
										case heartBeat = <-heartBeatChans[1]:
										case heartBeat = <-heartBeatChans[2]:
										case heartBeat = <-heartBeatChans[3]:
										case heartBeat = <-heartBeatChans[4]:
										case heartBeat = <-heartBeatChans[5]:
										case heartBeat = <-heartBeatChans[6]:
										case heartBeat = <-heartBeatChans[7]:
										case heartBeat = <-heartBeatChans[8]:
										case heartBeat = <-heartBeatChans[9]:
										case heartBeat = <-heartBeatChans[0]:
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
