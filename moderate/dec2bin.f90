program dec2bin
  character(32) :: arg
  character(10) :: str
  integer :: stat
  integer :: n
  call get_command_argument(1, arg)
  open(11, file=trim(arg), action='read')
  do
    read(11, '(a)', iostat=stat) str
    if (stat /= 0) exit
    if (str /= '') then
      read(str, '(i10)') i
      write(*,'(b0)') i
    end if
  end do
  close(11)
end program dec2bin

