docker rmi re-chinese-redis
docker rmi re-chinese-backend

# function green() { printf "\x1b[38;5;048m %s\e[0m " "${@}"; printf "\n"; }
# function red() { printf "\x1b[38;5;196m %s\e[0m " "${@}"; printf "\n"; }

# function good() {
#     cur_date="$(date -u +"${DATE_FMT}")"
#     local log_str="${cur_date} SUCCESS ${*}"
#     if [[ ${SILENT} -eq 0 ]]; then
#         green "${log_str}"
#     fi
# } # good - end

# function err() {
#     cur_date="$(date -u +"${DATE_FMT}")"
#     local log_str="${cur_date} ERROR ${*}"
#     red "${log_str}"
# } # err - end


# function run_main() {
#     var_last_status=$?
#     docker rmi re-chinese-redis
#     if [[ "${var_last_status}" == "0" ]]; then
#         err "Failed to remove image re-chinese-redis"
#     fi

#     var_last_status=$?
#     docker rmi re-chinese-backend
#     if [[ "${var_last_status}" == "0" ]]; then
#         err "Failed to remove image re-chinese-backend ${var_last_status}"
#     fi
# } # run_main - end

# run_main